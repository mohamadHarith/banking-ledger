package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mohamadHarith/banking-ledger/services/authentication-service/config"
	"github.com/mohamadHarith/banking-ledger/services/authentication-service/handler"
	"github.com/mohamadHarith/banking-ledger/services/authentication-service/repository"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/authentication_service_proto"
	"google.golang.org/grpc"
)

func main() {
	repo := repository.New()

	h := handler.New(repo)

	conf := config.GetConf()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.ServicePort))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterAuthServiceServer(srv, h)

	log.Printf("[%v] started on port [:%v]\n", conf.ServiceName, conf.ServicePort)

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
