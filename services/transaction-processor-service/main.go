package main

import (
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
	mq := mq.New()

	h := handler.New(repo, mq)

	srv := grpc.NewServer()
	pb.RegisterTransactionProcessorServiceServer(srv, h)

	log.Printf("[%v] started on port [:%v]\n", conf.ServiceName, conf.ServicePort)

	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

	// TODO: implement graceful shutdown
}
