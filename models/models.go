package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"iae.com/smartpower/config"
)

type User struct {
	gorm.Model
	Username string
	Email    string
}

type UserRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(config.DBDriver, config.GetDbConnectionString())

	if err != nil {
		fmt.Println(err)
		panic("fail to connect database")
	}

	db.AutoMigrate(&User{})
}
