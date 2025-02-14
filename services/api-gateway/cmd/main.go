package main

import (
	"context"
	"fmt"

	"github.com/mohamadHarith/banking-ledger/services/api-gateway/internal/mq"
)

func main() {
	msq := mq.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := make(chan uint32)
	go msq.ConsumeAccountBalance(ctx, res)

	for d := range res {
		fmt.Println(d)
	}

	select {}
}
