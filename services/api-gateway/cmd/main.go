package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/repository"
)

func main() {
	msq := mq.New()
	repo := repository.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan uint32)
	defer close(res)

	go msq.ConsumeAccountBalance(ctx, res)

	go func() {
		select {
		case <-ctx.Done():
			return
		case d := <-res:
			err := repo.SetUserBalance(ctx, "", d)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	mux := http.NewServeMux()
	log.Println("[api-gateway] started on port [:5002]")
	http.ListenAndServe(":5002", mux)
}
