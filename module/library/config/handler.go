package config

import (
	"net/http"
	"simple-library-app/module/library/internal/handler"

	"github.com/gorilla/mux"
)

func RegisterLibraryHandlers(router *mux.Router, config *LibraryConfig) {
	uc := NewLibraryUsecase(config)
	h := handler.NewLibraryHandler(uc)

	router.HandleFunc("/api/books", h.ListBook).Methods("GET")
	router.HandleFunc("/api/pickup-schedule/create", h.CreatePickupSchedule).Methods("POST")
	router.HandleFunc("/api/pickup-schedule", h.ListPickupSchedule).Methods("GET")

	router.Use(corsMiddleware)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
