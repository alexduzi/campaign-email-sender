package internalmock

import (
	"campaignemailsender/internal/contract"
	model "campaignemailsender/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) Get() ([]model.Campaign, error) {
	args := s.Called()
	return nil, args.Error(1)
}

func (s *CampaignServiceMock) GetByID(id string) (*contract.CampaignReduced, error) {
	args := s.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignReduced), nil
}

func (s *CampaignServiceMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CampaignServiceMock) Start(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CampaignServiceMock) SendEmailAndUpdateStatus(campaign *model.Campaign) {

}
