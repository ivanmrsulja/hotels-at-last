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
		{Name: "Hotel Lux", Address: "Address1", Base64Image: "Image"},
		{Name: "Hotel Vardar", Address: "Address2", Base64Image: "Image"},
		{Name: "Splendid", Address: "Address3", Base64Image: "Image"},
	}
	rooms = []model.Room{
		{NumberOfBeds: 2, Price: 100, AirConditioned: true, HasParkingSpace: true, HotelID: 1, HasTV: true},
		{NumberOfBeds: 4, Price: 200, AirConditioned: true, HasParkingSpace: true, HotelID: 1, HasTV: true},
		{NumberOfBeds: 4, Price: 130, AirConditioned: false, HasParkingSpace: false, HotelID: 1, HasTV: true},
		{NumberOfBeds: 3, Price: 90, AirConditioned: true, HasParkingSpace: false, HotelID: 2, HasTV: true},
		{NumberOfBeds: 3, Price: 80, AirConditioned: false, HasParkingSpace: false, HotelID: 2, HasTV: true},
		{NumberOfBeds: 5, Price: 420, AirConditioned: true, HasParkingSpace: false, HotelID: 3, HasTV: true},
		{NumberOfBeds: 5, Price: 530, AirConditioned: true, HasParkingSpace: true, HotelID: 3, HasTV: true},
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