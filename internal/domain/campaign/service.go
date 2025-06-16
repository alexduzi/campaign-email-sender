package campaign

import (
	"campaignemailsender/internal/contract"
	internalerrors "campaignemailsender/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImpl) Get() ([]Campaign, error) {
	result, err := s.Repository.Get()

	return result, err
}
