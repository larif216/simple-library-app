package handler_test

import (
	"net/http"
	"net/http/httptest"
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/handler"
	"simple-library-app/module/library/internal/usecase/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ListPickupScheduleHandlerTestSuite struct {
	suite.Suite
	mockLibraryUsecases *mock.MockLibraryUsecases
	handler             *handler.LibraryHandler
}

func (suite *ListPickupScheduleHandlerTestSuite) SetupTest() {
	suite.mockLibraryUsecases = new(mock.MockLibraryUsecases)
	suite.handler = handler.NewLibraryHandler(suite.mockLibraryUsecases)
}

func (suite *ListPickupScheduleHandlerTestSuite) TestListPickupSchedule_Success() {
	scheduleTime, _ := time.Parse(time.RFC3339, "2025-01-25T10:00:00Z")
	respPayload := []entity.PickupSchedule{
		{
			Book: entity.Book{
				EditionNumber: "12345",
				Title:         "Go Programming",
				Authors:       []string{"John Doe"},
				IsAvailable:   true,
			},
			DateTime: scheduleTime,
		},
	}

	suite.mockLibraryUsecases.On("ListPickupSchedule").Return(respPayload)

	request, err := http.NewRequest("GET", "/pickup-schedule", nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()

	suite.handler.ListPickupSchedule(rr, request)

	suite.Equal(http.StatusOK, rr.Code)

	expectedResponse := "[{\"ID\":0,\"Book\":{\"Title\":\"Go Programming\",\"Authors\":[\"John Doe\"],\"EditionNumber\":\"12345\",\"IsAvailable\":true},\"DateTime\":\"2025-01-25T10:00:00Z\"}]\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func TestListPickupScheduleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ListPickupScheduleHandlerTestSuite))
}
