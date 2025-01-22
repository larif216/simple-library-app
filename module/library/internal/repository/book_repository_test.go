package repository_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-library-app/module/library/internal/repository"
	umock "simple-library-app/module/library/internal/usecase/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookRepositoryTestSuite struct {
	suite.Suite
	repo         *repository.BookRepository
	mockClient   *umock.HTTPClient
	mockResponse *http.Response
}

func (suite *BookRepositoryTestSuite) SetupTest() {
	suite.mockClient = new(umock.HTTPClient)
	client := &http.Client{
		Transport: suite.mockClient,
	}
	suite.repo = repository.NewBookRepository("http://mockurl.com", client)

	suite.mockResponse = &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"name": "Test Subject", "works":[{"title":"Test Book 1", "authors":[{"name":"Author 1"},{"name":"Author 2"}], "availability":{"isbn":"12345"}}]}`))),
	}
}

func (suite *BookRepositoryTestSuite) TestGetBySubject() {
	suite.mockClient.On("RoundTrip", mock.Anything).Return(suite.mockResponse, nil)

	books, err := suite.repo.GetBySubject("test")

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), books, 1)
	assert.Equal(suite.T(), "Test Book 1", books[0].Title)
	assert.Contains(suite.T(), books[0].Authors, "Author 1")
	assert.Equal(suite.T(), "12345", books[0].EditionNumber)

	suite.mockClient.AssertExpectations(suite.T())
}

func (suite *BookRepositoryTestSuite) TestGetByEditionNumber() {
	suite.mockResponse.Body = ioutil.NopCloser(bytes.NewReader([]byte(`{"numFound": 1, "docs":[{"title":"Test Book 2", "author_name":["Author 1"]}]}`)))
	suite.mockClient.On("RoundTrip", mock.Anything).Return(suite.mockResponse, nil)

	book, err := suite.repo.GetByEditionNumber("12345")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Test Book 2", book.Title)
	assert.Contains(suite.T(), book.Authors, "Author 1")
	assert.Equal(suite.T(), "12345", book.EditionNumber)

	suite.mockClient.AssertExpectations(suite.T())
}

func (suite *BookRepositoryTestSuite) TestGetBySubject_FailedRequest() {
	mockErrorResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"error": "Internal Server Error"}`))),
	}
	suite.mockClient.On("RoundTrip", mock.Anything).Return(mockErrorResponse, nil)

	books, err := suite.repo.GetBySubject("test")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), books)
	assert.Equal(suite.T(), fmt.Sprintf("unexpected status code: %d", http.StatusInternalServerError), err.Error())

	suite.mockClient.AssertExpectations(suite.T())
}

func (suite *BookRepositoryTestSuite) TestGetByEditionNumber_FailedRequest() {
	mockErrorResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"error": "Internal Server Error"}`))),
	}
	suite.mockClient.On("RoundTrip", mock.Anything).Return(mockErrorResponse, nil)

	book, err := suite.repo.GetByEditionNumber("12345")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), book)
	assert.Equal(suite.T(), fmt.Sprintf("unexpected status code: %d", http.StatusInternalServerError), err.Error())

	suite.mockClient.AssertExpectations(suite.T())
}

func TestBookRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BookRepositoryTestSuite))
}
