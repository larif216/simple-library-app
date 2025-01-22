package usecase

import "simple-library-app/module/library/entity"

func (uc *LibraryUsecase) ListPickupSchedule() []entity.PickupSchedule {
	schedules := uc.pickupScheduleRepo.List()
	return schedules
}
