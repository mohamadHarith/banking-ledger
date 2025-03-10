package repository

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/config"
)

type Repository struct {
	db *sql.DB
}

var repo *Repository
var once sync.Once

func New() *Repository {
	once.Do(func() {
		cfg := config.GetConf()

		mysqlHost := cfg.MySql.ServiceName
		if cfg.IsLocalEnvironment() {
			mysqlHost = "localhost"
		}

		dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?multiStatements=true", cfg.MySql.User, cfg.MySql.Password, mysqlHost, cfg.MySql.Database)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		repo = &Repository{db: db}
	})

	if err := repo.migrateTables(); err != nil {
		panic(err)
	}

	return repo
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) migrateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id CHAR(36) PRIMARY KEY DEFAULT (uuid()),
		name VARCHAR(255) NOT NULL DEFAULT '',
		password VARCHAR(255) NOT NULL DEFAULT '',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS accounts (
		id CHAR(36) PRIMARY KEY DEFAULT (uuid()),
		user_id CHAR(36) NOT NULL,
		balance BIGINT UNSIGNED NOT NULL DEFAULT 0, -- stored in cents
		currency VARCHAR(10) DEFAULT 'MYR',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id CHAR(36)PRIMARY KEY DEFAULT (uuid()),
		account_id CHAR(36) NOT NULL,
		type ENUM('DEPOSIT', 'WITHDRAWAL', 'TRANSFER') NOT NULL,
		amount BIGINT NOT NULL, -- Stored in cents, withdrawals are negative values
		reference_id CHAR(36) NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE CASCADE
	);
	`

	_, err := r.db.Exec(query)

	return err
}
