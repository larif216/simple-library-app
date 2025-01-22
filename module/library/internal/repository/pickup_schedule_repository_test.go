package repository_test

import (
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PickupScheduleRepositoryTestSuite struct {
	suite.Suite
	repo *repository.PickupScheduleRepository
}

func (suite *PickupScheduleRepositoryTestSuite) SetupTest() {
	suite.repo = repository.NewPickupScheduleRepository()
}

func (suite *PickupScheduleRepositoryTestSuite) TestCreatePickupSchedule() {
	schedule := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Test Book",
			Authors:       []string{"Author 1"},
			EditionNumber: "12345",
		},
	}

	suite.repo.Create(schedule)

	assert.Equal(suite.T(), 1, schedule.ID)
	assert.Len(suite.T(), suite.repo.List(), 1)

	storedSchedule := suite.repo.List()[0]
	assert.Equal(suite.T(), "Test Book", storedSchedule.Book.Title)
	assert.Equal(suite.T(), "12345", storedSchedule.Book.EditionNumber)
	assert.Contains(suite.T(), storedSchedule.Book.Authors, "Author 1")
}

func (suite *PickupScheduleRepositoryTestSuite) TestCreatePickupSchedule_DuplicateEditionNumber() {
	schedule1 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 1",
			Authors:       []string{"Author 1"},
			EditionNumber: "12345",
		},
	}
	suite.repo.Create(schedule1)

	schedule2 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 2",
			Authors:       []string{"Author 2"},
			EditionNumber: "12345",
		},
	}
	suite.repo.Create(schedule2)

	assert.Equal(suite.T(), 2, schedule2.ID)
	assert.Len(suite.T(), suite.repo.List(), 2)

	schedules := suite.repo.List()
	assert.Equal(suite.T(), "Book 1", schedules[0].Book.Title)
	assert.Equal(suite.T(), "Book 2", schedules[1].Book.Title)
}

func (suite *PickupScheduleRepositoryTestSuite) TestGetByBookEditionNumber() {
	schedule1 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 1",
			Authors:       []string{"Author 1"},
			EditionNumber: "12345",
		},
	}
	schedule2 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 2",
			Authors:       []string{"Author 2"},
			EditionNumber: "67890",
		},
	}
	suite.repo.Create(schedule1)
	suite.repo.Create(schedule2)

	schedule := suite.repo.GetByBookEditionNumber("12345")
	assert.NotNil(suite.T(), schedule)
	assert.Equal(suite.T(), "Book 1", schedule.Book.Title)
	assert.Equal(suite.T(), "12345", schedule.Book.EditionNumber)

	schedule = suite.repo.GetByBookEditionNumber("00000")
	assert.Nil(suite.T(), schedule)
}

func (suite *PickupScheduleRepositoryTestSuite) TestGetByBookEditionNumber_EmptyRepository() {
	schedule := suite.repo.GetByBookEditionNumber("12345")
	assert.Nil(suite.T(), schedule)
}

func (suite *PickupScheduleRepositoryTestSuite) TestListPickupSchedules() {
	schedule1 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 1",
			Authors:       []string{"Author 1"},
			EditionNumber: "12345",
		},
	}
	schedule2 := &entity.PickupSchedule{
		Book: entity.Book{
			Title:         "Book 2",
			Authors:       []string{"Author 2"},
			EditionNumber: "67890",
		},
	}
	suite.repo.Create(schedule1)
	suite.repo.Create(schedule2)

	schedules := suite.repo.List()
	assert.Len(suite.T(), schedules, 2)
	assert.Equal(suite.T(), "Book 1", schedules[0].Book.Title)
	assert.Equal(suite.T(), "Book 2", schedules[1].Book.Title)
}

func (suite *PickupScheduleRepositoryTestSuite) TestListPickupSchedules_EmptyRepository() {
	schedules := suite.repo.List()
	assert.Len(suite.T(), schedules, 0)
}

func TestPickupScheduleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PickupScheduleRepositoryTestSuite))
}
