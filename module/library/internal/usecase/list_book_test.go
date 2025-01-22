package usecase_test

import (
	"errors"
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/usecase"
	umock "simple-library-app/module/library/internal/usecase/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ListBookUsecaseTestSuite struct {
	suite.Suite
	mockBookRepo   *umock.MockBookRepo
	mockPickupRepo *umock.MockPickupScheduleRepo
	libraryUsecase *usecase.LibraryUsecase
}

func (suite *ListBookUsecaseTestSuite) SetupTest() {
	suite.mockBookRepo = new(umock.MockBookRepo)
	suite.mockPickupRepo = new(umock.MockPickupScheduleRepo)
	suite.libraryUsecase = usecase.NewLibraryUseCase(suite.mockBookRepo, suite.mockPickupRepo)
}

func (suite *ListBookUsecaseTestSuite) TestListBook() {
	books := []*entity.Book{
		{
			Title:         "Book 1",
			Authors:       []string{"Author 1"},
			EditionNumber: "12345",
		},
		{
			Title:         "Book 2",
			Authors:       []string{"Author 2"},
			EditionNumber: "67890",
		},
		{
			Title:         "Book 3",
			Authors:       []string{"Author 3"},
			EditionNumber: "",
		},
	}

	suite.mockBookRepo.On("GetBySubject", "Science").Return(books, nil)

	suite.mockPickupRepo.On("GetByBookEditionNumber", "12345").Return(&entity.PickupSchedule{}).Once()
	suite.mockPickupRepo.On("GetByBookEditionNumber", "67890").Return(nil).Once()

	result, err := suite.libraryUsecase.ListBook("Science")

	assert.NoError(suite.T(), err)

	assert.Len(suite.T(), result, 2)

	assert.False(suite.T(), result[0].IsAvailable) // Book 1 should be unavailable
	assert.True(suite.T(), result[1].IsAvailable)  // Book 2 should be available
	assert.Equal(suite.T(), "Book 1", result[0].Title)
	assert.Equal(suite.T(), "Book 2", result[1].Title)
}

func (suite *ListBookUsecaseTestSuite) TestListBook_Error() {
	suite.mockBookRepo.On("GetBySubject", "Science").Return(nil, errors.New("repository error"))

	result, err := suite.libraryUsecase.ListBook("Science")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *ListBookUsecaseTestSuite) TestListBook_EmptySubject() {
	suite.mockBookRepo.On("GetBySubject", "NonExistentSubject").Return([]*entity.Book{}, nil)

	result, err := suite.libraryUsecase.ListBook("NonExistentSubject")

	assert.NoError(suite.T(), err)

	assert.Len(suite.T(), result, 0)
}

func TestListBookUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ListBookUsecaseTestSuite))
}
