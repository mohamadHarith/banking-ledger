package handler

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *transactionProcessorHandler) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (resp *pb.CreateAccountResponse, err error) {
	resp = &pb.CreateAccountResponse{}

	accountId := uuid.New().String()
	now := time.Now().In(time.Local)

	account := entity.Account{
		Id:        accountId,
		UserId:    *req.UserId,
		Balance:   *req.InitialBalance,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = h.repository.InsertAccount(ctx, &account)
	if err != nil {
		return nil, err
	}

	resp.Account = &pb.Account{
		Id:        &accountId,
		UserId:    &account.UserId,
		Balance:   &account.Balance,
		CreatedAt: timestamppb.New(account.CreatedAt),
		UpdatedAt: timestamppb.New(account.UpdatedAt),
	}

	go func() {
		err := h.mq.PublishAccountBalance(entity.Account{
			Id:      accountId,
			UserId:  *req.UserId,
			Balance: *req.InitialBalance,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		err := h.mq.PublishTransactionLog(entity.TransactionLog{
			Id:          uuid.NewString(),
			AccountId:   account.Id,
			UserId:      account.UserId,
			Balance:     *req.InitialBalance,
			Amount:      int32(account.Balance),
			Description: "Account Creation Initial Balance",
			CreatedAt:   account.CreatedAt,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return
}
