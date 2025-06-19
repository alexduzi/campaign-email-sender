package mock

import (
	"campaignemailsender/internal/contract"
	"campaignemailsender/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) Get() ([]campaign.Campaign, error) {
	args := s.Called()
	return nil, args.Error(1)
}

func (s *CampaignServiceMock) GetByID(id string) (*contract.CampaignReduced, error) {
	args := s.Called(id)
	return nil, args.Error(1)
}

func (s *CampaignServiceMock) Cancel(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
