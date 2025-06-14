package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	result, err := h.CampaignService.Repository.Get()

	return result, 200, err
}
