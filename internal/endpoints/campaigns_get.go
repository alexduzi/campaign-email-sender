package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	result, err := h.CampaignService.Get()

	return result, 200, err
	// return nil, 200, nil
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	result, err := h.CampaignService.GetByID(id)

	if err != nil {
		return nil, 400, err
	}

	return result, 200, err
}
