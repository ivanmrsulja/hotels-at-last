package utils

import (
	"fmt"
	"log"

	model "github.com/ivanmrsulja/hotels-at-last/review-service/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	reviews = []model.Review{
		{Comment: "Ovo je najjaci hotel nasvijet!", Rating: 4, TimesReported: 0, UserId: 3, RoomId: 1, UserUsername: "skafiskafnjak"},
		{Comment: "UZELI STE MI PARE MAJKU VAM [REDACTED]!", Rating: 1, TimesReported: 4, UserId: 3, RoomId: 2, UserUsername: "skafiskafnjak"},
		{Comment: "Bila je svadba i djani je izasao iz torte, sve nas je isprljao.", Rating: 2, TimesReported: 0, UserId: 2, RoomId: 1, UserUsername: "makulica"},
	}
)

var Db *gorm.DB
var err error

func ConnectToDatabase() {
	connectionString := "host=localhost user=postgres dbname=ReviewServiceDB sslmode=disable password=root port=5432"
	dialect := "postgres"

	Db, err = gorm.Open(dialect, connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to DB successfull.")
	}

	Db.DropTable("reviews")
	Db.AutoMigrate(&model.Review{})

	for _, review := range reviews {
		Db.Create(&review)
	}
}