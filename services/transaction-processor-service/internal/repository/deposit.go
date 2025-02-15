package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r *Repository) WithdrawOrDeposit(ctx context.Context, amount int32, userId, accountId string, now time.Time) (uint32, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return 0, err
	}

	query := `
	SELECT EXISTS (
		SELECT * FROM accounts WHERE user_id = ? AND id = ?
	);
	`

	var exists bool
	row := tx.QueryRowContext(ctx, query, userId, accountId)
	if err := row.Scan(&exists); err != nil {
		tx.Rollback()
		return 0, err
	}

	if !exists {
		return 0, fmt.Errorf("user account for %v does not exist", userId)
	}

	query = `
	SELECT balance FROM accounts WHERE user_id = ? AND id = ? FOR UPDATE;
	`

	var balance int32
	row = tx.QueryRowContext(ctx, query, userId, accountId)
	err = row.Scan(&balance)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	balance += amount
	if balance < 0 {
		return 0, fmt.Errorf("insufficient balance %v for amount %v", balance, amount)
	}

	query = `
	UPDATE accounts SET balance = ?, updated_at = ? WHERE user_id = ? AND id = ?;
	`

	_, err = tx.ExecContext(ctx, query, balance, now, userId, accountId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return uint32(balance), nil
}
