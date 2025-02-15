package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

type Handler struct {
	transactionProcessor pb.TransactionProcessorServiceClient
	validator            func(s interface{}) error
	repository           *repository.Repository
}

func New(r *repository.Repository) *Handler {

	conn, err := grpc.NewClient("localhost:5001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewTransactionProcessorServiceClient(conn)

	return &Handler{
		transactionProcessor: client,
		validator:            validator.New(validator.WithRequiredStructEnabled()).Struct,
		repository:           r,
	}

}

func writeResp(w http.ResponseWriter, message string, statusCode int32, item any, items any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))
	msg := dto.ResponseMessage{
		ErrorCode: statusCode,
		Item:      item,
		Items:     items,
		Message:   message,
	}
	j, _ := json.Marshal(msg)
	fmt.Fprint(w, string(j))
}
