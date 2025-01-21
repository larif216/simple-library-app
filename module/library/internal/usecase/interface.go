package usecase

import "simple-library-app/module/library/entity"

type BookRepository interface {
	GetBySubject(subject string) ([]*entity.Book, error)
}

type LibraryUsecases interface {
	ListBook(subject string) ([]*entity.Book, error)
}
