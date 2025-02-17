package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

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
	transactionProcessor  pb.TransactionProcessorServiceClient
	transactionProcessors []pb.TransactionProcessorServiceClient
	transactionLogger     pb3.TransactionLoggerServiceClient
	authenticator         pb2.AuthServiceClient
	validator             func(s interface{}) error
	repository            *repository.Repository
}

var handler *Handler
var once sync.Once

func New(r *repository.Repository) *Handler {
	once.Do(func() {

		conf := config.GetConf()

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

		h := &Handler{
			transactionLogger: client3,
			validator:         validator.New(validator.WithRequiredStructEnabled()).Struct,
			repository:        r,
			authenticator:     client2,
		}

		// multiple transaction processors for load balancing
		if conf.IsDevelopmentEnvironment() {
			for i := 1; i <= 2; i++ {
				transactionProcessorHost := conf.TransactionProcessorService.ServiceName + strconv.Itoa(i)

				conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", transactionProcessorHost, conf.TransactionProcessorService.ServicePort), grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					panic(err)
				}

				client := pb.NewTransactionProcessorServiceClient(conn)
				h.transactionProcessors = append(h.transactionProcessors, client)
			}
		} else {
			transactionProcessorHost := "localhost"

			conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", transactionProcessorHost, conf.TransactionProcessorService.ServicePort), grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				panic(err)
			}

			client := pb.NewTransactionProcessorServiceClient(conn)
			h.transactionProcessor = client
		}

		handler = h
	})

	return handler
}

var currentIndex uint32

func (h *Handler) getNextTransactionProcessor() pb.TransactionProcessorServiceClient {
	conf := config.GetConf()
	if conf.IsLocalEnvironment() {
		return h.transactionProcessor
	}

	index := atomic.AddUint32(&currentIndex, 1)
	return h.transactionProcessors[(index-1)%uint32(len(h.transactionProcessors))]
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
