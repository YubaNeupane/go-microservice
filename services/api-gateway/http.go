package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "BAD_REQUEST",
			Message: "Failed to parse request body",
		})
		return
	}

	log.Default().Println("Preview Trip Request:", reqBody)

	defer r.Body.Close()

	if reqBody.UserID == "" {
		writeJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "BAD_REQUEST",
			Message: "userID is required",
		})
		return
	}
	//TODO: Call Trip Service to get trip preview details
	response := &contracts.APIResponse{
		Error: nil,
		Data:  "OK",
	}

	writeJSON(w, http.StatusCreated, response)
}

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip Start Endpoint"))
}
