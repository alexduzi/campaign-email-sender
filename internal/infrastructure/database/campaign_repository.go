package database

import (
	"campaignemailsender/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return errors.New("an error")
}
