package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"
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

	response := &contracts.APIResponse{
		Error: nil,
		Data:  "OK",
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)

	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, &contracts.APIError{
			Code:    "FAILED TO PARSE",
			Message: "Failed to parse response data from trip service",
		})
		return
	}

	response.Data = respBody

	util.WriteJSON(w, http.StatusCreated, response)
}

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Trip Start Endpoint"))
}
