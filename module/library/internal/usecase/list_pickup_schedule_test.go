package usecase_test

import (
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/usecase"
	"simple-library-app/module/library/internal/usecase/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ListPickupScheduleUsecaseTestSuite struct {
	suite.Suite
	mockPickupScheduleRepo *mock.MockPickupScheduleRepo
	uc                     *usecase.LibraryUsecase
}

func (suite *ListPickupScheduleUsecaseTestSuite) SetupTest() {
	suite.mockPickupScheduleRepo = new(mock.MockPickupScheduleRepo)
	suite.uc = usecase.NewLibraryUseCase(nil, suite.mockPickupScheduleRepo)
}

func (suite *ListPickupScheduleUsecaseTestSuite) TestListPickupSchedule() {
	mockSchedules := []entity.PickupSchedule{
		{Book: entity.Book{EditionNumber: "12345"}, DateTime: time.Now()},
		{Book: entity.Book{EditionNumber: "67890"}, DateTime: time.Now()},
	}

	suite.mockPickupScheduleRepo.On("List").Return(mockSchedules)

	result := suite.uc.ListPickupSchedule()

	suite.NotNil(result)
	suite.Len(result, len(mockSchedules))
	suite.Equal(mockSchedules, result)

	suite.mockPickupScheduleRepo.AssertExpectations(suite.T())
}

func TestListPickupScheduleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ListPickupScheduleUsecaseTestSuite))
}
