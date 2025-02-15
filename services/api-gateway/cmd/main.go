package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/handler"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
)

func main() {
	msq := mq.New()
	repo := repository.New()

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
				log.Println("account balance => ", r)
				err := repo.SetUserBalance(ctx, r.UserId, r.Id, r.Balance)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	mux := http.NewServeMux()

	h := handler.New(repo)

	mux.Handle("/account", http.HandlerFunc(h.CreateAccount))
	// // mux.Handle("/login", nil)
	mux.Handle("/deposit", http.HandlerFunc(h.Deposit))
	mux.Handle("/withdraw", http.HandlerFunc(h.Withdraw))
	// mux.Handle("/transfer", nil)
	mux.Handle("/balance", http.HandlerFunc(h.GetBalance))

	log.Println("[api-gateway] started on port [:5002]")
	http.ListenAndServe(":5002", mux)
}
