package config

import (
	"simple-library-app/module/library/internal/repository"
	"simple-library-app/module/library/internal/usecase"
)

func NewLibraryUsecase(cfg *LibraryConfig) *usecase.LibraryUsecase {
	bookRepo := repository.NewBookRepository(cfg.BaseURL, cfg.HTTPClient)

	uc := usecase.NewLibraryUseCase(
		bookRepo,
	)

	return uc
}
