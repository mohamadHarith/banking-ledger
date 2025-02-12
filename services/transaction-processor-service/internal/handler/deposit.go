package handler

import (
	"context"

	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

func (h *transactionProcessorHandler) Deposit(ctx context.Context, req *pb.DepositRequest) (resp *pb.DepositResponse, err error) {
	resp = &pb.DepositResponse{}

	err = h.repository.Deposit(ctx, *req.Amount, *req.UserId, *req.AccountId)
	if err != nil {
		return resp, err
	}

	return
}
