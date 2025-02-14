package repository

import "context"

func (r *Repository) InsertAccount(ctx context.Context, accountId, userId string, initialAmount uint32) error {

	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts (id, user_id, balance) VALUES (?, ?, ?)",
		accountId,
		userId,
		initialAmount,
	)

	return err
}
