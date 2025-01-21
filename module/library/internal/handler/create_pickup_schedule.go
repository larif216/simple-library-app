package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-library-app/internal/util"
	"simple-library-app/module/library/entity"
)

func (h *LibraryHandler) CreatePickupSchedule(w http.ResponseWriter, r *http.Request) {
	var req entity.CreatePickupScheduleRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.WriteHTTPResponse(w, map[string]string{
			"message": "Invalid request body. Must be valid JSON.",
		}, http.StatusBadRequest)
		return
	}

	if req.EditionNumber == "" || req.DateTime == "" {
		util.WriteHTTPResponse(w, map[string]string{
			"message": "Missing or invalid parameters: edition_number or datetime",
		}, http.StatusBadRequest)
		return
	}

	resp, err := h.libraryUsecase.CreatePickupSchedule(req)
	if err != nil {
		util.WriteHTTPResponse(w, map[string]string{
			"message": fmt.Sprintf("Error creating pickup schedule: %v", err),
		}, http.StatusInternalServerError)
		return
	}

	util.WriteHTTPResponse(w, resp, http.StatusOK)
}
