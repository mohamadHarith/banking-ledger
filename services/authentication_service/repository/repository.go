package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohamadHarith/banking-ledger/services/authentication_service/config"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	cfg := config.GetConf()

	dsn := fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v?multiStatements=true", cfg.MySql.User, cfg.MySql.Password, cfg.MySql.Database)

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
	CREATE TABLE IF NOT EXISTS users (
		id CHAR(36) PRIMARY KEY DEFAULT (uuid()),
		full_name VARCHAR(255) NOT NULL DEFAULT '',
		user_name VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL DEFAULT '',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
	`

	_, err := r.db.Exec(query)

	return err
}

func (r *Repository) CreateUser(ctx context.Context, u entity.User) error {

	query := `INSERT INTO users (id, full_name, user_name, password, created_at, updated_at) VALUE (?,?,?,?,?,?)`

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUser(ctx context.Context, id string) (u entity.User, err error) {

	query := `SELECT * FROM users where id = ?`

	row := r.db.QueryRowContext(ctx, query)

	err = row.Scan(&u.Id, &u.FullName, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return
	}

	return
}
