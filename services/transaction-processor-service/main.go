package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/config"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

func main() {

	conf := config.GetConf()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.ServicePort))
	if err != nil {
		panic(err)
	}

	repo := repository.New()
	defer repo.Close()

	mq := mq.New()
	defer mq.Close()

	h := handler.New(repo, mq)

	opt := grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log.Printf("Received request - Method: %s - Request: %+v", info.FullMethod, req)

		return handler(ctx, req)
	})

	srv := grpc.NewServer(opt)
	pb.RegisterTransactionProcessorServiceServer(srv, h)

	log.Printf("[%v] started on port [:%v]\n", conf.ServiceName, conf.ServicePort)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

	// TODO: implement graceful shutdown
}
