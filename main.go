package main

import (
	db "api-station/database"
	"api-station/route"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_db := db.DBCon()
	route.SetupRoute(_db)

}