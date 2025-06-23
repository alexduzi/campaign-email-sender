package main

import (
	"campaignemailsender/internal/infrastructure/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	println("Started Campaign Worker")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when loading the .env file")
	}

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}
	campaigns, err := repository.GetCampaignsToBeSent()

	if err != nil {
		println(err.Error())
	}

	for _, camp := range campaigns {
		println(camp.ID)
	}
}
