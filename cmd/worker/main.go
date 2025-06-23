package main

import (
	"campaignemailsender/internal/domain/campaign"
	"campaignemailsender/internal/infrastructure/database"
	mailer "campaignemailsender/internal/infrastructure/mail"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Started Campaign Worker...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when loading the .env file")
	}

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}
	service := campaign.ServiceImpl{
		Repository: &database.CampaignRepository{Db: db},
		SendEmail:  mailer.SendEmail,
	}

	log.Println("Getting campaigns to send email...")

	var campaigns []campaign.Campaign

	for {
		campaigns, err = repository.GetCampaignsToBeSent()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("Found %v campaigns...\n", len(campaigns))

		if len(campaigns) == 0 {
			time.Sleep(time.Minute * 5)
			continue
		}

		for _, camp := range campaigns {
			log.Printf("Sending email for campaign %v - %v ", camp.ID, camp.Name)
			service.SendEmailAndUpdateStatus(&camp)
		}

		log.Println("Done...")
		time.Sleep(time.Minute * 30)
	}
}
