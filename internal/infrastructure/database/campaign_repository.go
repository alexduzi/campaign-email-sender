package database

import (
	model "campaignemailsender/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Create(campaign *model.Campaign) error {
	tx := c.Db.Create(campaign)

	return tx.Error
}

func (c *CampaignRepository) Get() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	tx := c.Db.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetByID(id string) (*model.Campaign, error) {
	var campaign model.Campaign
	tx := c.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}

func (c *CampaignRepository) Update(campaign *model.Campaign) error {
	tx := c.Db.Save(campaign)

	return tx.Error
}

func (c *CampaignRepository) Delete(campaign *model.Campaign) error {
	c.Db.Select("Contacts").Delete(campaign)

	tx := c.Db.Delete(campaign)

	return tx.Error
}

func (c *CampaignRepository) GetCampaignsToBeSent() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	tx := c.Db.Preload("Contacts").Find(&campaigns,
		"status = ? and date_part('minute', now()::timestamp - updated_on::timestamp) > ?",
		model.Started,
		1)
	return campaigns, tx.Error
}
