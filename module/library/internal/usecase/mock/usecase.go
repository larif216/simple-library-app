package mock

import (
	"simple-library-app/module/library/entity"

	"github.com/stretchr/testify/mock"
)

type MockLibraryUsecases struct {
	mock.Mock
}

func (m *MockLibraryUsecases) ListBook(subject string) ([]*entity.Book, error) {
	args := m.Called(subject)
	if books, ok := args.Get(0).([]*entity.Book); ok {
		return books, args.Error(1)
	}

	return []*entity.Book{}, args.Error(1)
}

func (m *MockLibraryUsecases) CreatePickupSchedule(req entity.CreatePickupScheduleRequest) (*entity.CreatePickupScheduleResponse, error) {
	args := m.Called(req)
	if resp, ok := args.Get(0).(*entity.CreatePickupScheduleResponse); ok {
		return resp, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockLibraryUsecases) ListPickupSchedule() []entity.PickupSchedule {
	args := m.Called()
	return args.Get(0).([]entity.PickupSchedule)
}
