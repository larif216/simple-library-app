package handler

import "simple-library-app/module/library/internal/usecase"

type LibraryHandler struct {
	libraryUsecase usecase.LibraryUsecases
}

func NewLibraryHandler(uc usecase.LibraryUsecases) *LibraryHandler {
	return &LibraryHandler{
		libraryUsecase: uc,
	}
}
