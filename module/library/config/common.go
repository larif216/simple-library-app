package config

import (
	"simple-library-app/module/library/internal/repository"
	"simple-library-app/module/library/internal/usecase"
)

func NewLibraryUsecase(cfg *LibraryConfig) *usecase.LibraryUsecase {
	bookRepo := repository.NewBookRepository(cfg.BaseURL, cfg.HTTPClient)
	pickupScheduleRepo := repository.NewPickupScheduleRepository()

	uc := usecase.NewLibraryUseCase(
		bookRepo,
		pickupScheduleRepo,
	)

	return uc
}
