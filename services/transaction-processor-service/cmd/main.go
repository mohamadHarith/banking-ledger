package main

import (
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/transaction-processor-service/internal/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_processor_proto"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		panic(err)
	}

	repo := repository.New()

	srv := grpc.NewServer()
	pb.RegisterTransactionProcessorServiceServer(srv, handler.New(repo))

	log.Println("transaction-processor-service started on port 5001")

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

	// TODO: implement graceful shutdown
}
