package campaign_test

import (
	"campaignemailsender/internal/contract"
	"campaignemailsender/internal/domain/campaign"
	internalerrors "campaignemailsender/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	internalmock "campaignemailsender/internal/test/internal-mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test V",
		Content: "Body hi!",
		Emails: []string{
			"teste1@test.com",
		},
		CreatedBy: "admin@admin.com",
	}

	service = campaign.ServiceImpl{}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Create_CreateCampaign(t *testing.T) {

	repositoryMock := new(internalmock.RepositoryMock)

	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDataBaseSave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.RepositoryMock)

	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_GetById_Return_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	repositoryMock := new(internalmock.RepositoryMock)

	repositoryMock.On("GetByID", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	service.Repository = repositoryMock

	campaignReturned, _ := service.GetByID(campaign.ID)

	assert.NotNil(campaign)
	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
	assert.Equal(campaign.CreatedBy, campaignReturned.CreatedBy)
}

func Test_GetById_Return_Error(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	repositoryMock := new(internalmock.RepositoryMock)

	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("error"))

	service.Repository = repositoryMock

	campaignReturned, err := service.GetByID(campaign.ID)

	assert.Nil(campaignReturned)
	assert.NotNil(err)
}

func Test_Delete_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Delete("invalid_id")

	assert.ErrorIs(err, internalerrors.NotFound)
}

func Test_Delete_CampaignStatusInvalid_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)
	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test 1", "Body Test!", []string{"test@test.com.br"}, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(errors.New("error to delete campaign"))
	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_when_delete_success(t *testing.T) {
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test 1", "Body Test!", []string{"test@test.com.br"}, newCampaign.CreatedBy)
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(nil)
	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}

func Test_Start_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)
	campaignIdInvalid := "invalid_id"
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Start(campaignIdInvalid)

	assert.Equal(internalerrors.NotFound.Error(), err.Error())
}

func Test_Start_CampaignStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	assert := assert.New(t)
	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock

	err := service.Start(campaign.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_Start_should_send_mail(t *testing.T) {
	assert := assert.New(t)
	campaignMock := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	// mocking SendEmail func
	sentMail := false
	sendMail := func(c *campaign.Campaign) error {
		if c.ID == campaignMock.ID {
			sentMail = true
		}
		return nil
	}
	service.SendEmail = sendMail

	service.Start(campaignMock.ID)

	assert.True(sentMail)
}

func Test_Start_should_send_mail_throw_error(t *testing.T) {
	assert := assert.New(t)
	campaignMock := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	// mocking SendEmail func
	sendMail := func(c *campaign.Campaign) error {
		return errors.New("error to send mail")
	}
	service.SendEmail = sendMail

	errMail := service.Start(campaignMock.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), errMail.Error())
}

func Test_Start_return_nil_when_updated_to_done(t *testing.T) {
	assert := assert.New(t)
	campaignMock := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock := new(internalmock.RepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(campaignMock, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return campaignMock.ID == c.ID && c.Status == campaign.Done
	})).Return(nil)
	service.Repository = repositoryMock

	// mocking SendEmail func
	sendMail := func(c *campaign.Campaign) error {
		return nil
	}
	service.SendEmail = sendMail

	service.Start(campaignMock.ID)

	assert.Equal(campaign.Done, campaignMock.Status)
}
