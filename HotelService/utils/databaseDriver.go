package utils

import (
	"fmt"
	"log"

	model "github.com/ivanmrsulja/hotels-at-last/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	hotels = []model.Hotel{
		{Name: "Hotel Lux"},
		{Name: "Hotel Vardar"},
		{Name: "Splendid"},
	}
	rooms = []model.Room{
		{NumberOfBeds: 2, AirConditioned: true, HotelID: 1},
		{NumberOfBeds: 4, AirConditioned: true, HotelID: 1},
		{NumberOfBeds: 4, AirConditioned: false, HotelID: 1},
		{NumberOfBeds: 3, AirConditioned: true, HotelID: 2},
		{NumberOfBeds: 3, AirConditioned: false, HotelID: 2},
		{NumberOfBeds: 5, AirConditioned: true, HotelID: 3},
		{NumberOfBeds: 5, AirConditioned: true, HotelID: 3},
	}
)

var Db *gorm.DB
var err error

func ConnectToDatabase() {
	connectionString := "host=localhost user=postgres dbname=HotelServiceDB sslmode=disable password=root port=5432"
	dialect := "postgres"

	Db, err = gorm.Open(dialect, connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to DB successfull.")
	}

	Db.DropTable("rooms")
	Db.DropTable("hotels")
	Db.AutoMigrate(&model.Hotel{})
	Db.AutoMigrate(&model.Room{})

	for _, hotel := range hotels {
		Db.Create(&hotel)
	}
	for _, room := range rooms {
		Db.Create(&room)
	}
}