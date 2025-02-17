package handler

import (
	"context"
	"sync"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionProcessorHandler struct {
	pb.UnimplementedTransactionProcessorServiceServer
	repository *repository.Repository
	mq         *mq.MQ
}

var handler *transactionProcessorHandler
var once sync.Once

func New(repo *repository.Repository, mq *mq.MQ) *transactionProcessorHandler {
	once.Do(func() {
		handler = &transactionProcessorHandler{
			repository: repo,
			mq:         mq,
		}
	})

	return handler
}

func (h *transactionProcessorHandler) Transfer(ctx context.Context, req *pb.TransferRequest) (resp *emptypb.Empty, err error) {
	return
}
