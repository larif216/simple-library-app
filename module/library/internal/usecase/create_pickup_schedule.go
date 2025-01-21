package usecase

import (
	"errors"
	"simple-library-app/module/library/entity"
	"time"
)

func (uc *LibraryUsecase) CreatePickupSchedule(req entity.CreatePickupScheduleRequest) (*entity.CreatePickupScheduleResponse, error) {
	parsedDateTime, err := time.Parse(time.RFC3339, req.DateTime)
	if err != nil {
		return nil, errors.New("invalid DateTime format. Must be in RFC3339 format (e.g., 2025-01-25T10:00:00Z)")
	}

	existSchedule := uc.pickupScheduleRepo.GetByBookEditionNumber(req.EditionNumber)
	if existSchedule != nil {
		return nil, errors.New("book have been scheduled for pick up")
	}

	book, err := uc.bookRepo.GetByEditionNumber(req.EditionNumber)
	if err != nil {
		return nil, err
	}

	schedule := &entity.PickupSchedule{
		Book:     *book,
		DateTime: parsedDateTime,
	}

	uc.pickupScheduleRepo.Create(schedule)

	return &entity.CreatePickupScheduleResponse{
		Schedule: schedule,
		Message:  "Pickup schedule created!",
	}, nil
}
