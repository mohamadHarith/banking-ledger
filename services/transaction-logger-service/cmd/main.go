package main

import (
	"context"
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
	"google.golang.org/grpc"
)

func main() {

	repo := repository.New()
	mq := mq.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan entity.TransactionLog)
	defer close(res)

	go mq.ConsumeTransactionLog(ctx, res)

	go func() {
		for txn := range res {
			log.Println(txn)
			if err := repo.InsertTransactionLog(ctx, txn); err != nil {
				log.Println(err)
			}
		}
	}()

	srv := grpc.NewServer()
	h := handler.New(repo)
	pb.RegisterTransactionLoggerServiceServer(srv, h)

	lis, err := net.Listen("tcp", "localhost:5004")
	if err != nil {
		panic(err)
	}

	log.Println("started at port 5004")

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
