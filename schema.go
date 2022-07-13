package main

// import (
// 	"time"
// )

type User struct {
	UserId       int64
	FirstName    string
	LastName     string
	Email        string
	Contact      int64
	HashPassword string
	Gender       string
	CreatedAt    string
}

type Color struct {
	CategoryId int64
	Color      string
	Category   string
}

type Activity struct {
	ActivitiesId int64
	Category     string
	CategoryId   int64
	Activity     string
	UserId       int64
	Color        string
	Status       string
	CreatedAt    string
}

// type Tabler interface {
// 	TableName() string
// }

// TableName overrides the table name used by User to `profiles`
// func (Activity) TableName() string {
// 	return "activity"
// }
