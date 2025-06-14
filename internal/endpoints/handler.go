package endpoints

import "campaignemailsender/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
