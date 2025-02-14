package main

import (
	"context"
	"log"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/mq"
	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/repository"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
)

func main() {

	repo := repository.New()
	mq := mq.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan entity.TransactionLog)
	defer close(res)

	if err := mq.ConsumeTransactionLog(ctx, res); err != nil {
		panic(err)
	}

	go func() {
		for txn := range res {
			if err := repo.InsertTransactionLog(ctx, txn); err != nil {
				log.Println(err)
			}
		}
	}()

	select {}

}
