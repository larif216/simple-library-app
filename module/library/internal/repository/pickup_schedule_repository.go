package repository

import (
	"simple-library-app/module/library/entity"
	"sync"
)

type PickupScheduleRepository struct {
	mutex           sync.Mutex
	pickupSchedules []entity.PickupSchedule
	nextID          int
}

func NewPickupScheduleRepository() *PickupScheduleRepository {
	return &PickupScheduleRepository{
		pickupSchedules: make([]entity.PickupSchedule, 0),
		nextID:          1,
	}
}

func (r *PickupScheduleRepository) Create(schedule entity.PickupSchedule) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	schedule.ID = r.nextID
	r.nextID++
	r.pickupSchedules = append(r.pickupSchedules, schedule)
}

func (r *PickupScheduleRepository) FindByBookEditionNumber(editionNumber string) *entity.Book {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, schedule := range r.pickupSchedules {
		if schedule.Book.EditionNumber == editionNumber {
			return &schedule.Book
		}
	}

	return nil
}
