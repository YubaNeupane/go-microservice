package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/types"
	"ride-sharing/shared/util"
)

type HttpHandler struct {
	Sevice domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandlePreview(w http.ResponseWriter, r *http.Request) {

	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "BAD_REQUEST",
			Message: "Failed to parse request body",
		})
		return
	}

	// fare := &domain.RideFareModel{
	// 	UserID:            reqBody.UserID,
	// 	TotalPriceInCents: 25.50,
	// }
	ctx := r.Context()

	t, err := h.Sevice.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Println("Created Trip")
	}

	util.WriteJSON(w, http.StatusCreated, t)
}
