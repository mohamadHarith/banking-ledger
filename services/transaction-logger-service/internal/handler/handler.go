package handler

import (
	"context"
	"log"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedTransactionLoggerServiceServer
	repository *repository.Repository
}

func New(r *repository.Repository) *Handler {
	return &Handler{
		repository: r,
	}
}

func (h *Handler) GetTransactionLogs(ctx context.Context, req *pb.GetTransactionLogsRequest) (resp *pb.GetTransactionLogsResponse, err error) {
	resp = &pb.GetTransactionLogsResponse{}

	txnLogs, totNoOfRecords, totNoOfPages, hasNext, hasPrev, err := h.repository.GetTransactionLogs(ctx, *req.Page, *req.UserId, *req.AccountId)
	if err != nil {
		return nil, err
	}

	log.Println("here=> ", txnLogs)

	resp.CurrentPage = req.Page
	resp.TotalPages = &totNoOfPages
	resp.TotalRecords = &totNoOfRecords

	nextPage := uint32(0)
	prevPage := uint32(0)

	if hasNext {
		nextPage = *req.Page + 1
	}

	if hasPrev {
		prevPage = *req.Page - 1
	}

	resp.NextPage = &nextPage
	resp.PrevPage = &prevPage

	resp.TransactionLogs = make([]*pb.TransactionLog, 0)
	for i := 0; i < len(txnLogs); i++ {
		txnLog := txnLogs[i]
		resp.TransactionLogs = append(resp.TransactionLogs, &pb.TransactionLog{
			Id:          &txnLog.Id,
			AccountId:   &txnLog.AccountId,
			UserId:      &txnLog.UserId,
			Amount:      &txnLog.Amount,
			Balance:     &txnLog.Balance,
			Description: &txnLog.Description,
			CreatedAt:   timestamppb.New(txnLog.CreatedAt),
		})
	}

	return
}
