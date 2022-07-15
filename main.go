package main

import (
	"fmt"

	"github.com/dotjarr/activity-service/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	//password.go
	// password := "secret"
	// hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	// fmt.Println("Password:", password)
	// fmt.Println("Hash:    ", hash)

	// match := CheckPasswordHash(password, hash)
	// fmt.Println("Match:   ", match)
	///////////////////////////////////////////////////////////////////////////
	//	  dsn :="user=postgres password=postgres1234 dbname=dotjar sslmode=disable"
	dsn := "postgres://postgres:postgres1234@localhost/dotjar?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err.Error())
	}
	//	defer db.Close()
	// var activity Activity
	database, err := db.DB()
	err = database.Ping()

	if err != nil {
		panic(err.Error())

	}
	fmt.Println("Connection to Postgresql was successful!")

	// _ = db.Take(&activity)
	// //	result.RowsAffected
	// //	result.Error
	// fmt.Println(activity.Activity)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080

	// r.GET("/allActivity", func(k *gin.Context) {
	// 	var activities []Activity

	// 	_ = db.Joins("JOIN users ON users.user_id = activities.user_id").
	// 		Where("users.first_name=? ", "arvind").
	// 		Find(&activities)

	// 	k.JSON(200, gin.H{

	// 		"message": activities,
	// 	})
	// })

	r.POST("/signup", handlers.Signup(db))

	r.Run()

}
