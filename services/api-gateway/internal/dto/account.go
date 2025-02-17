package dto

import (
	"time"

	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

type Account struct {
	Id        string    `json:"id"`
	UserId    string    `json:"userId"`
	Balance   uint32    `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromProtoToEntityAccount(i *pb.Account) (o entity.Account) {
	o.Id = *i.Id
	o.UserId = *i.UserId
	o.Balance = *i.Balance
	o.CreatedAt = i.CreatedAt.AsTime()
	o.UpdatedAt = i.UpdatedAt.AsTime()
	return
}

func FromEntityToDtoAccount(i *entity.Account) (o Account) {
	o.Id = i.Id
	o.UserId = i.UserId
	o.Balance = i.Balance

	o.CreatedAt = i.CreatedAt
	o.UpdatedAt = i.UpdatedAt
	return
}
