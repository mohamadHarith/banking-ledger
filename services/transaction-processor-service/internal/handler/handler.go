package handler

import (
	"context"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
)

type transactionProcessorHandler struct {
	pb.UnimplementedTransactionProcessorServiceServer
	repository *repository.Repository
}

func New(repo *repository.Repository) *transactionProcessorHandler {
	return &transactionProcessorHandler{
		repository: repo,
	}
}

func (h *transactionProcessorHandler) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (resp *pb.WithdrawResponse, err error) {
	return
}

func (h *transactionProcessorHandler) Transfer(ctx context.Context, req *pb.TransferRequest) (resp *pb.TransferResponse, err error) {
	return
}
