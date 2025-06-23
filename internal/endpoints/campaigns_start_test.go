package endpoints

import (
	internalerrors "campaignemailsender/internal/internal-errors"
	internalmock "campaignemailsender/internal/test/internal-mock"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_should_return_204(t *testing.T) {
	assert := assert.New(t)

	campaignId := "valid_id"

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("POST", "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(204, status)
	assert.Nil(err)
}

func Test_CampaignStart_should_return_500(t *testing.T) {
	assert := assert.New(t)

	campaignId := "valid_id"

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(internalerrors.ErrInternal)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("POST", "/", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignId)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	res := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(res, req)

	assert.Equal(500, status)
	assert.NotNil(err)
	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
