package entity

import "time"

type TransactionLog struct {
	Id          string    `json:"id" bson:"_id"`
	AccountId   string    `json:"accountId" bson:"accountId"`
	UserId      string    `json:"userId" bson:"userId"`
	Amount      int32     `json:"amount" bson:"amount"`
	Balance     uint32    `json:"balance" bson:"balance"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}
