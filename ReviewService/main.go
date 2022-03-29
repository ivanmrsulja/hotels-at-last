package main

import (
	router "github.com/ivanmrsulja/hotels-at-last/review-service/router"
	util "github.com/ivanmrsulja/hotels-at-last/review-service/utils"
)

func main() {
	util.ConnectToDatabase()
	router.HandleRequests()

	defer util.Db.Close()
}