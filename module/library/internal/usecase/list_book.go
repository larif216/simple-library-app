package usecase

import "simple-library-app/module/library/internal/entity"

func (uc *LibraryUsecase) ListBook(subject string) ([]*entity.Book, error) {
	books, err := uc.bookRepo.GetBySubject(subject)
	if err != nil {
		return nil, err
	}

	return books, nil
}
