package handler

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *transactionProcessorHandler) Deposit(ctx context.Context, req *pb.DepositRequest) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}

	now := time.Now().In(time.Local)

	balance, err := h.repository.WithdrawOrDeposit(ctx, int32(*req.Amount), *req.UserId, *req.AccountId, now)
	if err != nil {
		return resp, err
	}

	description := "Cash Deposit"
	if req.Description != nil && *req.Description != "" {
		description = *req.Description
	}

	go func() {
		err := h.mq.PublishAccountBalance(entity.Account{
			UserId:  *req.UserId,
			Id:      *req.AccountId,
			Balance: balance,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		err := h.mq.PublishTransactionLog(entity.TransactionLog{
			Id:          uuid.NewString(),
			AccountId:   *req.AccountId,
			UserId:      *req.UserId,
			Amount:      int32(*req.Amount),
			Balance:     balance,
			Description: description,
			CreatedAt:   now,
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return
}
