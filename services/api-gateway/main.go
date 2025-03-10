package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/config"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
)

func main() {
	msq := mq.New()
	defer msq.Close()

	repo := repository.New()
	defer repo.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan entity.Account)
	defer close(res)

	go msq.ConsumeAccountBalance(ctx, res)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case r := <-res:
				err := repo.SetUserBalance(ctx, r.UserId, r.Id, r.Balance)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	mux := http.NewServeMux()

	h := handler.New(repo)

	mux.Handle("/user", http.HandlerFunc(h.CreateUser))
	mux.Handle("/login", http.HandlerFunc(h.Login))

	mux.Handle("/account", h.Authenticate(http.HandlerFunc(h.CreateAccount)))
	mux.Handle("/deposit", h.Authenticate(http.HandlerFunc(h.Deposit)))
	mux.Handle("/withdraw", h.Authenticate(http.HandlerFunc(h.Withdraw)))
	mux.Handle("/balance", h.Authenticate(http.HandlerFunc(h.GetBalance)))
	mux.Handle("/transactions", h.Authenticate(h.GetTransactions))

	conf := config.GetConf()

	log.Printf("[%v] started on port [:%v]\n", conf.ServiceName, conf.ServicePort)

	err := http.ListenAndServe(fmt.Sprintf(":%v", conf.ServicePort), mux)
	if err != nil {
		panic(err)
	}
}
