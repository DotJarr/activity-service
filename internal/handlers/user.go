package handlers

import (
	"github.com/dotjarr/activity-service/pkg/user"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"time"
)

type Request struct {
	UserId       int64
	FirstName    string
	LastName     string
	Email        string
	Contact      int64
	HashPassword string
	Gender       string
	CreatedAt    time.Time
}

type Response struct {
	Id uint
}

func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(500, "validation error")
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.HashPassword), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, "error")

			return
		}

		newUser := &user.User{
			// UserId:       req.UserId,
			FirstName:    req.FirstName,
			LastName:     req.LastName,
			Email:        req.Email,
			Contact:      req.Contact,
			HashPassword: string(passwordHash),
			Gender:       req.Gender,
			CreatedAt:    req.CreatedAt,
		}

		_ = db.Omit("UserId", "CreatedAt").Create(&newUser)
		c.JSON(200, "successful")
	}

}
