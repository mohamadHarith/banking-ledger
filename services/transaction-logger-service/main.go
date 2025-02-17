package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/config"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
	"google.golang.org/grpc"
)

func main() {

	repo := repository.New()
	defer repo.Close()

	mq := mq.New()
	defer mq.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan entity.TransactionLog)
	defer close(res)

	go mq.ConsumeTransactionLog(ctx, res)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case txn := <-res:
				log.Println(txn)
				if err := repo.InsertTransactionLog(ctx, txn); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	srv := grpc.NewServer()
	h := handler.New(repo)
	pb.RegisterTransactionLoggerServiceServer(srv, h)

	conf := config.GetConf()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.ServicePort))
	if err != nil {
		panic(err)
	}

	log.Printf("[%v] started on port [:%v]\n", conf.ServiceName, conf.ServicePort)

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
