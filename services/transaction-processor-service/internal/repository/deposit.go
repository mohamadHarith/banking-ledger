package repository

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *Repository) Deposit(ctx context.Context, amount uint32, userId, accountId string) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `
	SELECT EXISTS (
		SELECT * FROM ACCOUNTS WHERE user_id = $1 AND id = $2;
	);
	`

	var exists bool
	row := tx.QueryRowContext(ctx, query, userId, accountId)
	if err := row.Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("user account for %v does not exist", userId)
	}

	query = `
	SELECT balance FROM accounts WHERE user_id = $1 AND id = $2 FOR UPDATE;
	`

	var balance uint32
	row = tx.QueryRowContext(ctx, query, userId, accountId)
	err = row.Scan(&balance)
	if err != nil {
		return err
	}

	balance += amount

	query = `
	UPDATE accounts SET balance = $1 WHERE user_id $2 AND id = $3;
	`

	_, err = tx.ExecContext(ctx, query, balance, userId, accountId)
	if err != nil {
		return err
	}

	query = `INSERT INTO accounts columns(user_id, balance) VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, query, userId, accountId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
