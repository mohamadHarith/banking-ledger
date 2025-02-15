package entity

import "time"

type Account struct {
	Id        string
	UserId    string
	Balance   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
