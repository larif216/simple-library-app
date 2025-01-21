package handler

import (
	"fmt"
	"net/http"
	"simple-library-app/internal/util"
)

func (h *LibraryHandler) ListBook(w http.ResponseWriter, r *http.Request) {
	subject := r.URL.Query().Get("subject")
	if subject == "" {
		util.WriteHTTPResponse(w, map[string]string{
			"message": "Invalid 'subject' parameter",
		}, http.StatusBadRequest)
		return
	}

	resp, err := h.libraryUsecase.ListBook(subject)
	if err != nil {
		util.WriteHTTPResponse(w, map[string]string{
			"message": fmt.Sprintf("Error getting books: %v", err),
		}, http.StatusInternalServerError)
		return
	}

	util.WriteHTTPResponse(w, resp, http.StatusOK)
}
