package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-library-app/module/library/entity"
	"simple-library-app/module/library/internal/handler"
	"simple-library-app/module/library/internal/usecase/mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListBookHandlerTestSuite struct {
	suite.Suite
	mockLibraryUsecases *mock.MockLibraryUsecases
	handler             *handler.LibraryHandler
}

func (suite *ListBookHandlerTestSuite) SetupTest() {
	suite.mockLibraryUsecases = new(mock.MockLibraryUsecases)
	suite.handler = handler.NewLibraryHandler(suite.mockLibraryUsecases)
}

func (suite *ListBookHandlerTestSuite) TestListBook_Success() {
	books := []*entity.Book{
		{EditionNumber: "12345", Title: "Go Programming", Authors: []string{"John Doe"}, IsAvailable: true},
		{EditionNumber: "67890", Title: "Advanced Go", Authors: []string{"Jane Doe"}, IsAvailable: true},
	}

	suite.mockLibraryUsecases.On("ListBook", "go").Return(books, nil)

	req, err := http.NewRequest("GET", "/books?subject=go", nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.ListBook(rr, req)

	suite.Equal(http.StatusOK, rr.Code)

	expectedResponse := "[{\"Title\":\"Go Programming\",\"Authors\":[\"John Doe\"],\"EditionNumber\":\"12345\",\"IsAvailable\":true},{\"Title\":\"Advanced Go\",\"Authors\":[\"Jane Doe\"],\"EditionNumber\":\"67890\",\"IsAvailable\":true}]\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func (suite *ListBookHandlerTestSuite) TestListBook_InvalidSubject() {
	req, err := http.NewRequest("GET", "/books", nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.ListBook(rr, req)

	suite.Equal(http.StatusBadRequest, rr.Code)
	expectedResponse := "{\"message\":\"Invalid 'subject' parameter\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func (suite *ListBookHandlerTestSuite) TestListBook_InternalServerError() {
	suite.mockLibraryUsecases.On("ListBook", "go").Return(nil, errors.New("internal server error"))

	req, err := http.NewRequest("GET", "/books?subject=go", nil)
	suite.NoError(err)

	rr := httptest.NewRecorder()
	suite.handler.ListBook(rr, req)

	suite.Equal(http.StatusInternalServerError, rr.Code)
	expectedResponse := "{\"message\":\"Error getting books: internal server error\"}\n"
	suite.Equal(expectedResponse, rr.Body.String())
}

func TestListBookHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ListBookHandlerTestSuite))
}
