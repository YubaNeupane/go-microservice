package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"
	"strconv"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "BAD_REQUEST",
			Message: "Failed to parse request body",
		})
		return
	}

	log.Default().Println("Preview Trip Request:", reqBody)

	defer r.Body.Close()

	if reqBody.UserID == "" {
		util.WriteJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "BAD_REQUEST",
			Message: "userID is required",
		})
		return
	}
	//TODO: Call Trip Service to get trip preview details

	tripService, err := grpc_clients.NewTripServiceClient()

	if err != nil {
		log.Fatal(err)
		response := &contracts.APIResponse{
			Error: &contracts.APIError{
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Failed to preview trip",
			},
		}

		util.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}
	defer tripService.Close()

	previewTripResp, err := tripService.Client.PreviewTrip(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("Failed to preview a trip %v", err)
		return
	}

	response := &contracts.APIResponse{
		Data: previewTripResp,
	}

	util.WriteJSON(w, http.StatusCreated, response)
}

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip Start Endpoint"))
}
