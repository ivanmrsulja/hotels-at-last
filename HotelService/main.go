package main

import (
	router "github.com/ivanmrsulja/hotels-at-last/router"
	util "github.com/ivanmrsulja/hotels-at-last/utils"
)

func main() {
	util.ConnectToDatabase()
	router.HandleRequests()

	defer util.Db.Close()
}