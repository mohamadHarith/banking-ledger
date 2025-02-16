package main

import (
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/authentication_service/handler"
	"github.com/mohamadHarith/banking-ledger/services/authentication_service/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.New()

	h := handler.New(repo)

	lis, err := net.Listen("tcp", "localhost:5005")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterAuthServiceServer(srv, h)

	log.Println("starting authentication server on port 5005")

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}

}
