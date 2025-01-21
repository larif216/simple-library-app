package config

import (
	"simple-library-app/module/library/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterLibraryHandlers(router *mux.Router, config *LibraryConfig) {
	uc := NewLibraryUsecase(config)
	h := handler.NewLibraryHandler(uc)

	router.HandleFunc("/api/books", h.ListBook).Methods("GET")
	router.HandleFunc("/api/pickup-schedule/create", h.CreatePickupSchedule).Methods("POST")
}
