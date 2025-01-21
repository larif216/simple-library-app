package usecase

type LibraryUsecase struct {
	bookRepo BookRepository
}

func NewLibraryUseCase(
	bookRepo BookRepository,
) *LibraryUsecase {
	return &LibraryUsecase{
		bookRepo: bookRepo,
	}
}
