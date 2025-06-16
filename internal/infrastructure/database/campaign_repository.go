package database

import (
	"campaignemailsender/internal/domain/campaign"
	"fmt"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}

func (c *CampaignRepository) GetByID(id string) (*campaign.Campaign, error) {
	for _, campaign := range c.campaigns {
		if campaign.ID == id {
			return &campaign, nil
		}
	}
	return nil, fmt.Errorf("campaign with id: %s not found", id)
}
