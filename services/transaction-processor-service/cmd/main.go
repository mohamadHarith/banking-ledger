package main

import (
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		panic(err)
	}

	repo := repository.New()
	mq := mq.New()
	// mq.PublishAccountBalance(123)
	// time.Sleep(time.Second * 2)
	// mq.PublishAccountBalance(124)
	// time.Sleep(time.Second * 2)

	// mq.PublishAccountBalance(125)

	h := handler.New(repo, mq)

	srv := grpc.NewServer()
	pb.RegisterTransactionProcessorServiceServer(srv, h)

	log.Println("transaction-processor-service started on port 5001")

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

	// TODO: implement graceful shutdown
}
