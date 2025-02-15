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

func (r *PickupScheduleRepository) Create(schedule *entity.PickupSchedule) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	schedule.ID = r.nextID
	r.nextID++
	r.pickupSchedules = append(r.pickupSchedules, *schedule)
}

func (r *PickupScheduleRepository) GetByBookEditionNumber(editionNumber string) *entity.PickupSchedule {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, schedule := range r.pickupSchedules {
		if schedule.Book.EditionNumber == editionNumber {
			return &schedule
		}
	}

	return nil
}

func (r *PickupScheduleRepository) List() []entity.PickupSchedule {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	schedulesCopy := make([]entity.PickupSchedule, len(r.pickupSchedules))
	copy(schedulesCopy, r.pickupSchedules)

	return schedulesCopy
}
