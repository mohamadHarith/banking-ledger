package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/dto"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
	pb2 "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
	pb3 "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

type Handler struct {
	transactionProcessor pb.TransactionProcessorServiceClient
	transactionLogger    pb3.TransactionLoggerServiceClient
	authenticator        pb2.AuthServiceClient
	validator            func(s interface{}) error
	repository           *repository.Repository
}

var handler *Handler
var once sync.Once

func New(r *repository.Repository) *Handler {
	once.Do(func() {

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

		transactionLoggerHost := conf.TransactionLoggerService.ServiceName
		if conf.IsLocalEnvironment() {
			transactionLoggerHost = "localhost"
		}

		conn3, err := grpc.NewClient(fmt.Sprintf("%v:%v", transactionLoggerHost, conf.TransactionLoggerService.ServicePort), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			panic(err)
		}

		client3 := pb3.NewTransactionLoggerServiceClient(conn3)

		handler = &Handler{
			transactionProcessor: client,
			transactionLogger:    client3,
			validator:            validator.New(validator.WithRequiredStructEnabled()).Struct,
			repository:           r,
			authenticator:        client2,
		}
	})

	return handler
}

func writeResp(w http.ResponseWriter, message string, statusCode int32, item any, items []any, pagination ...dto.Pagination) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))
	msg := dto.ResponseMessage{
		ErrorCode: statusCode,
		Item:      item,
		Items:     items,
		Message:   message,
	}
	if len(pagination) > 0 {
		msg.Pagination = &pagination[0]
	}
	j, _ := json.Marshal(msg)
	fmt.Fprint(w, string(j))
}
