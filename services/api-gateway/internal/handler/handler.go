package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
	pb2 "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

type Handler struct {
	transactionProcessor pb.TransactionProcessorServiceClient
	authenticator        pb2.AuthServiceClient
	validator            func(s interface{}) error
	repository           *repository.Repository
}

func New(r *repository.Repository) *Handler {

	conf := config.GetConf()

	transactionProcessorHost := conf.TransactionProcessorService.ServiceName
	if conf.IsLocalEnvironment() {
		transactionProcessorHost = "localhost"
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", transactionProcessorHost, conf.TransactionProcessorService.ServicePort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewTransactionProcessorServiceClient(conn)

	authenticationServiceHost := conf.AuthenticationService.ServiceName
	if conf.IsLocalEnvironment() {
		authenticationServiceHost = "localhost"
	}

	conn2, err := grpc.NewClient(fmt.Sprintf("%v:%v", authenticationServiceHost, conf.AuthenticationService.ServicePort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client2 := pb2.NewAuthServiceClient(conn2)

	return &Handler{
		transactionProcessor: client,
		validator:            validator.New(validator.WithRequiredStructEnabled()).Struct,
		repository:           r,
		authenticator:        client2,
	}
}

func writeResp(w http.ResponseWriter, message string, statusCode int32, item any, items []any) {
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
