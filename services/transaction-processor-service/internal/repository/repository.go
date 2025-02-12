package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/config"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	cfg := config.GetConf()

	dsn := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v", cfg.MySql.User, cfg.MySql.Password, cfg.MySql.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	r := &Repository{db: db}

	if err := r.migrateTables(); err != nil {
		panic(err)
	}

	return r
}

func (r *Repository) migrateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
    	id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    	user_id CHAR(36) NOT NULL,
    	balance BIGINT UNSIGNED NOT NULL DEFAULT 0, -- stored in cents
    	currency VARCHAR(10) DEFAULT 'USD',
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS transactions (
    	id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    	account_id CHAR(36) NOT NULL,
    	type ENUM('DEPOSIT', 'WITHDRAWAL', 'TRANSFER') NOT NULL,
    	amount BIGINT NOT NULL, -- Stored in cents, withdrawals are negative values
    	reference_id CHAR(36) NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
	);
	`

	_, err := r.db.Exec(query)

	return err
}
