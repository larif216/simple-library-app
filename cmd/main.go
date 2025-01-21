package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"simple-library-app/internal/config"
	"syscall"
	"time"
)

func main() {
	httpServer, err := config.NewHttpServer()
	if err != nil {
		log.Fatalf("Error initializing the HTTP server: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Server starting on %s", httpServer.Config.Host)

	go func() {
		if err := httpServer.HTTPServer.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop

	log.Println("Shutting down the server...")

	shutdownTimeout := 5 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := httpServer.HTTPServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server stopped")
}
