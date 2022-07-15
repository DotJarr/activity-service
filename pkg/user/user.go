package user

import (
	"time"
)

type User struct {
	UserId       int64
	FirstName    string
	LastName     string
	Email        string
	Contact      int64
	HashPassword string
	Gender       string
	CreatedAt    time.Time
}
