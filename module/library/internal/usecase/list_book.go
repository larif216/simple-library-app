package usecase

import (
	"simple-library-app/module/library/entity"
)

func (uc *LibraryUsecase) ListBook(subject string) ([]*entity.Book, error) {
	books, err := uc.bookRepo.GetBySubject(subject)
	if err != nil {
		return nil, err
	}

	var filteredBooks []*entity.Book
	for _, book := range books {
		if book.EditionNumber == "" {
			continue
		}
		schedule := uc.pickupScheduleRepo.GetByBookEditionNumber(book.EditionNumber)

		book.IsAvailable = schedule == nil

		filteredBooks = append(filteredBooks, book)
	}

	return filteredBooks, nil
}
