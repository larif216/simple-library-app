package usecase

import "simple-library-app/module/library/entity"

type BookRepository interface {
	GetBySubject(subject string) ([]*entity.Book, error)
	GetByEditionNumber(editionNumber string) (*entity.Book, error)
}

type PickupScheduleRepository interface {
	Create(schedule *entity.PickupSchedule)
	GetByBookEditionNumber(editionNumber string) *entity.PickupSchedule
}

type LibraryUsecases interface {
	ListBook(subject string) ([]*entity.Book, error)
	CreatePickupSchedule(req entity.CreatePickupScheduleRequest) (*entity.CreatePickupScheduleResponse, error)
}
