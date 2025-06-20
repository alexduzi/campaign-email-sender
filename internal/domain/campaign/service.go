package campaign

import (
	"campaignemailsender/internal/contract"

	// "campaignemailsender/internal/domain/campaign"
	internalerrors "campaignemailsender/internal/internal-errors"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]Campaign, error)
	GetByID(id string) (*contract.CampaignReduced, error)
	Delete(id string) error
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImpl) Get() ([]Campaign, error) {
	result, err := s.Repository.Get()

	return result, err
}

func (s *ServiceImpl) GetByID(id string) (*contract.CampaignReduced, error) {
	result, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, internalerrors.GetError(err)
	}

	return &contract.CampaignReduced{
		ID:                   result.ID,
		Name:                 result.Name,
		Content:              result.Content,
		Status:               result.Status,
		AmountOfEmailsToSend: len(result.Contacts),
	}, nil
}

func (s *ServiceImpl) Delete(id string) error {
	result, err := s.Repository.GetByID(id)
	if err != nil {
		return internalerrors.GetError(err)
	}

	if result.Status != Pending {
		return errors.New("campaign status invalid")
	}

	result.Delete()

	errSave := s.Repository.Delete(result)

	if errSave != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
