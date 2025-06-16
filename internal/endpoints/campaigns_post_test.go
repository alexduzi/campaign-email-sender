package endpoints

import (
	"bytes"
	"campaignemailsender/internal/contract"
	"campaignemailsender/internal/domain/campaign"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *serviceMock) Get() ([]campaign.Campaign, error) {
	args := r.Called()
	return nil, args.Error(1)
}

func (r *serviceMock) GetByID(id string) (*contract.CampaignReduced, error) {
	args := r.Called(id)
	return nil, args.Error(1)
}

var (
	body = contract.NewCampaign{
		Name:    "test",
		Content: "Hi everyone",
		Emails:  []string{"test@test.com"},
	}
)

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.Content == body.Content
	})).Return("1234567", nil)

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)

	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
