package handler

import (
	"net/http"
	"simple-library-app/internal/util"
)

func (h *LibraryHandler) ListPickupSchedule(w http.ResponseWriter, r *http.Request) {
	resp := h.libraryUsecase.ListPickupSchedule()

	util.WriteHTTPResponse(w, resp, http.StatusOK)
}
