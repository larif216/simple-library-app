package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/handler"
	"simple-library-app/module/library/internal/usecase/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type CreatePickupScheduleHandlerTestSuite struct {
	suite.Suite
	mockLibraryUsecases *mock.MockLibraryUsecases
	handler             *handler.LibraryHandler
}

func (suite *CreatePickupScheduleHandlerTestSuite) SetupTest() {
	suite.mockLibraryUsecases = new(mock.MockLibraryUsecases)
	suite.handler = handler.NewLibraryHandler(suite.mockLibraryUsecases)
}

func (suite *CreatePickupScheduleHandlerTestSuite) TestCreatePickupSchedule_Success() {
	scheduleTime, err := time.Parse(time.RFC3339, "2025-01-25T10:00:00Z")
	suite.NoError(err)

	reqPayload := `{
		"edition_number": "12345",
		"datetime": "2025-01-25T10:00:00Z"
	}`

	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "2025-01-25T10:00:00Z",
	}

	respPayload := &entity.CreatePickupScheduleResponse{
		Schedule: &entity.PickupSchedule{
			Book:     entity.Book{EditionNumber: "12345", Title: "Go Programming", Authors: []string{"John Doe"}, IsAvailable: true},
			DateTime: scheduleTime,
		},
		Message: "Pickup schedule created!",
	}

	suite.mockLibraryUsecases.On("CreatePickupSchedule", req).Return(respPayload, nil)

	reqBody := bytes.NewBufferString(reqPayload)
	request, err := http.NewRequest("POST", "/pickup-schedule", reqBody)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.CreatePickupSchedule(rr, request)

	suite.Equal(http.StatusOK, rr.Code)

	expectedResponse := "{\"Schedule\":{\"ID\":0,\"Book\":{\"Title\":\"Go Programming\",\"Authors\":[\"John Doe\"],\"EditionNumber\":\"12345\",\"IsAvailable\":true},\"DateTime\":\"2025-01-25T10:00:00Z\"},\"Message\":\"Pickup schedule created!\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func (suite *CreatePickupScheduleHandlerTestSuite) TestCreatePickupSchedule_InvalidJSON() {
	reqPayload := `{
		"edition_number": "12345",
		"datetime": "2025-01-25T10:00:00Z"
	` // Missing closing brace

	reqBody := bytes.NewBufferString(reqPayload)
	request, err := http.NewRequest("POST", "/pickup-schedule", reqBody)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.CreatePickupSchedule(rr, request)

	suite.Equal(http.StatusBadRequest, rr.Code)

	expectedResponse := "{\"message\":\"Invalid request body. Must be valid JSON.\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func (suite *CreatePickupScheduleHandlerTestSuite) TestCreatePickupSchedule_MissingParameters() {
	reqPayload := `{
		"edition_number": "12345"
	}` // Missing DateTime parameter

	reqBody := bytes.NewBufferString(reqPayload)
	request, err := http.NewRequest("POST", "/pickup-schedule", reqBody)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.CreatePickupSchedule(rr, request)

	suite.Equal(http.StatusBadRequest, rr.Code)

	expectedResponse := "{\"message\":\"Missing or invalid parameters: edition_number or datetime\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func (suite *CreatePickupScheduleHandlerTestSuite) TestCreatePickupSchedule_InternalServerError() {
	reqPayload := `{
		"edition_number": "12345",
		"datetime": "2025-01-25T10:00:00Z"
	}`

	req := entity.CreatePickupScheduleRequest{
		EditionNumber: "12345",
		DateTime:      "2025-01-25T10:00:00Z",
	}

	suite.mockLibraryUsecases.On("CreatePickupSchedule", req).Return(nil, errors.New("internal server error"))

	reqBody := bytes.NewBufferString(reqPayload)
	request, err := http.NewRequest("POST", "/pickup-schedule", reqBody)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.CreatePickupSchedule(rr, request)

	suite.Equal(http.StatusInternalServerError, rr.Code)

	expectedResponse := "{\"message\":\"Error creating pickup schedule: internal server error\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func TestCreatePickupScheduleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreatePickupScheduleHandlerTestSuite))
}
