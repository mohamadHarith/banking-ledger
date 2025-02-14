package entity

type TransactionLog struct {
	Amount      uint32 `json:"amount" bson:"amount"`
	Description string `json:"description" bson:"description"`
}
