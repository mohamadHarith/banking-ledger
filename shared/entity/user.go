package entity

import "time"

type User struct {
	Id        string
	FullName  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
