package shared

import "github.com/mohamadHarith/banking-ledger/shared/entity"

const MQExchange = "transaction_exchange"
const TransactionBalanceRoutingKey = "transaction.balance"
const TransactionLogRoutingKey = "transaction.log"
const TransactionBalanceQueue = "transaction_balance_queue"
const TransactionLogQueue = "transaction_log_queue"

type AccountBalanceMessage struct {
	Balance uint32 `json:"balance"`
}

type TransactionLogMessage struct {
	entity.TransactionLog
}
