package utils

import (
	"fmt"
	"log"
	"time"

	model "github.com/ivanmrsulja/hotels-at-last/user-service/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	users = []model.User{
		{Email: "email1@email.com", Password: "password1", Name: "Ivica", Surname: "Roganovic", Role: model.ADMIN, BannedUntil: time.Now()},
		{Email: "email2@email.com", Password: "password2", Name: "Jovan", Surname: "Mustur", Role: model.USER, BannedUntil: time.Now()},
		{Email: "email3@email.com", Password: "password3", Name: "Marko", Surname: "Kapisoda", Role: model.USER, BannedUntil: time.Now().Add(time.Minute * 30)},
	}
)

var Db *gorm.DB
var err error

func ConnectToDatabase() {
	connectionString := "host=localhost user=postgres dbname=UserServiceDB sslmode=disable password=root port=5432"
	dialect := "postgres"

	Db, err = gorm.Open(dialect, connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to DB successfull.")
	}

	Db.DropTable("users")
	Db.AutoMigrate(&model.User{})

	for _, user := range users {
		Db.Create(&user)
	}
}