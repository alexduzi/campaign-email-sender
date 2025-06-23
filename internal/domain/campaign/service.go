package campaign

import (
	"campaignemailsender/internal/contract"

	// "campaignemailsender/internal/domain/campaign"
	internalerrors "campaignemailsender/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]Campaign, error)
	GetByID(id string) (*contract.CampaignReduced, error)
	Delete(id string) error
	Start(id string) error
}

type ServiceImpl struct {
	Repository Repository
	SendEmail  func(campaign *Campaign) error
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

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
		return nil, internalerrors.GetError(err, "")
	}

	return &contract.CampaignReduced{
		ID:                   result.ID,
		Name:                 result.Name,
		Content:              result.Content,
		Status:               result.Status,
		AmountOfEmailsToSend: len(result.Contacts),
		CreatedBy:            result.CreatedBy,
	}, nil
}

func (s *ServiceImpl) Delete(id string) error {
	result, err := s.getCampaignWithStatusValidation(id)
	if err != nil {
		return err
	}

	result.Delete()

	errSave := s.Repository.Delete(result)

	if errSave != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) Start(id string) error {
	result, err := s.getCampaignWithStatusValidation(id)
	if err != nil {
		return err
	}

	go func() {
		err := s.SendEmail(result)
		if err != nil {
			result.Fail()
		} else {
			result.Done()
		}
		s.Repository.Update(result)
	}()

	result.Started()

	err = s.Repository.Update(result)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) getCampaignWithStatusValidation(id string) (*Campaign, error) {
	result, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, internalerrors.GetError(err, "")
	}

	if result.Status != Pending {
		return nil, internalerrors.GetError(err, "campaign status invalid")
	}

	return result, nil
}
