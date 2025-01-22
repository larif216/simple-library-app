package usecase_test

import (
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/usecase"
	umock "simple-library-app/module/library/internal/usecase/mock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreatePickupScheduleUsecaseTestSuite struct {
	suite.Suite
	mockPickupScheduleRepo *umock.MockPickupScheduleRepo
	mockBookRepo           *umock.MockBookRepo
	uc                     *usecase.LibraryUsecase
}

func (suite *CreatePickupScheduleUsecaseTestSuite) SetupTest() {
	suite.mockPickupScheduleRepo = new(umock.MockPickupScheduleRepo)
	suite.mockBookRepo = new(umock.MockBookRepo)
	suite.uc = usecase.NewLibraryUseCase(suite.mockBookRepo, suite.mockPickupScheduleRepo)
}

func (suite *CreatePickupScheduleUsecaseTestSuite) TestCreatePickupSchedule() {
	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "2025-01-25T10:00:00Z",
	}

	book := &entity.Book{
		EditionNumber: "12345",
		Title:         "Test Book",
	}

	suite.mockPickupScheduleRepo.On("GetByBookEditionNumber", req.EditionNumber).Return(nil)
	suite.mockBookRepo.On("GetByEditionNumber", req.EditionNumber).Return(book, nil)
	suite.mockPickupScheduleRepo.On("Create", mock.Anything).Return()

	resp, err := suite.uc.CreatePickupSchedule(req)

	suite.NoError(err)
	suite.NotNil(resp)
	suite.Equal("Pickup schedule created!", resp.Message)

	suite.mockPickupScheduleRepo.AssertExpectations(suite.T())
	suite.mockBookRepo.AssertExpectations(suite.T())
}

func (suite *CreatePickupScheduleUsecaseTestSuite) TestCreatePickupSchedule_InvalidDateTimeFormat() {
	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "invalid-date-time",
	}

	resp, err := suite.uc.CreatePickupSchedule(req)

	suite.Error(err)
	suite.Nil(resp)
	suite.Equal("invalid DateTime format. Must be in RFC3339 format (e.g., 2025-01-25T10:00:00Z)", err.Error())

	suite.mockPickupScheduleRepo.AssertExpectations(suite.T())
	suite.mockBookRepo.AssertExpectations(suite.T())
}

func (suite *CreatePickupScheduleUsecaseTestSuite) TestCreatePickupSchedule_BookAlreadyScheduled() {
	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "2025-01-25T10:00:00Z",
	}

	existingSchedule := &entity.PickupSchedule{}
	suite.mockPickupScheduleRepo.On("GetByBookEditionNumber", req.EditionNumber).Return(existingSchedule)

	resp, err := suite.uc.CreatePickupSchedule(req)

	suite.Error(err)
	suite.Nil(resp)
	suite.Equal("book have been scheduled for pick up", err.Error())

	suite.mockPickupScheduleRepo.AssertExpectations(suite.T())
	suite.mockBookRepo.AssertExpectations(suite.T())
}

func (suite *CreatePickupScheduleUsecaseTestSuite) TestCreatePickupSchedule_BookNotFound() {
	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "2025-01-25T10:00:00Z",
	}

	suite.mockPickupScheduleRepo.On("GetByBookEditionNumber", req.EditionNumber).Return(nil)
	suite.mockBookRepo.On("GetByEditionNumber", req.EditionNumber).Return(nil, nil)

	resp, err := suite.uc.CreatePickupSchedule(req)

	suite.Error(err)
	suite.Nil(resp)
	suite.Equal("book not found", err.Error())

	suite.mockPickupScheduleRepo.AssertNotCalled(suite.T(), "Create", mock.Anything)
	suite.mockPickupScheduleRepo.AssertExpectations(suite.T())
	suite.mockBookRepo.AssertExpectations(suite.T())
}

func TestCreatePickupScheduleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreatePickupScheduleUsecaseTestSuite))
}
