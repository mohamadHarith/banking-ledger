package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

func (h *transactionProcessorHandler) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (resp *pb.CreateAccountResponse, err error) {
	resp = &pb.CreateAccountResponse{}

	accountId := uuid.New().String()

	err = h.repository.InsertAccount(ctx, accountId, *req.UserId, *req.InitialBalance)

	resp.AccountId = &accountId

	// FIXME: publish by user id and account id
	go h.mq.PublishAccountBalance(*req.InitialBalance)

	return
}
