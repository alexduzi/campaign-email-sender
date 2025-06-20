package endpoints

import (
	"campaignemailsender/internal/contract"
	internalmock "campaignemailsender/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGet_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignResponse := contract.CampaignReduced{
		ID:        "343",
		Name:      "Hi everyone",
		Content:   "Hi everyone",
		Status:    "Pending",
		CreatedBy: "admin@admin.com",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetByID", mock.Anything).Return(&campaignResponse, nil)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.GetByID(res, req)

	assert.Equal(200, status)
	assert.Equal(campaignResponse.ID, response.(*contract.CampaignReduced).ID)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignReduced).Name)
	assert.Equal(campaignResponse.Content, response.(*contract.CampaignReduced).Content)
	assert.Equal(campaignResponse.Status, response.(*contract.CampaignReduced).Status)
	assert.Equal(campaignResponse.CreatedBy, response.(*contract.CampaignReduced).CreatedBy)
}

func Test_CampaignGet_should_return_error_when_campaign_dont_exists(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("GetByID", mock.Anything).Return(nil, errors.New("error"))

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, status, err := handler.GetByID(res, req)

	assert.Equal(400, status)
	assert.NotNil(err)
}
