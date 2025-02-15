package repository

import (
	"context"

	"github.com/mohamadHarith/banking-ledger/shared/entity"
)

func (r *Repository) InsertAccount(ctx context.Context, account *entity.Account) error {

	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts (id, user_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		account.Id,
		account.UserId,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	return err
}
