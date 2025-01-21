package usecase

type LibraryUsecase struct {
	bookRepo           BookRepository
	pickupScheduleRepo PickupScheduleRepository
}

func NewLibraryUseCase(
	bookRepo BookRepository,
	pickupScheduleRepo PickupScheduleRepository,
) *LibraryUsecase {
	return &LibraryUsecase{
		bookRepo:           bookRepo,
		pickupScheduleRepo: pickupScheduleRepo,
	}
}
