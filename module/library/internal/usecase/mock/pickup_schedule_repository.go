package mock

import (
	"simple-library-app/module/library/entity"

	"github.com/stretchr/testify/mock"
)

// MockPickupScheduleRepo is a mock for the PickupScheduleRepository interface
type MockPickupScheduleRepo struct {
	mock.Mock
}

func (m *MockPickupScheduleRepo) Create(schedule *entity.PickupSchedule) {
	m.Called(schedule)
}

func (m *MockPickupScheduleRepo) GetByBookEditionNumber(editionNumber string) *entity.PickupSchedule {
	args := m.Called(editionNumber)
	if schedule, ok := args.Get(0).(*entity.PickupSchedule); ok {
		return schedule
	}

	return nil
}

func (m *MockPickupScheduleRepo) List() []entity.PickupSchedule {
	args := m.Called()
	return args.Get(0).([]entity.PickupSchedule)
}
