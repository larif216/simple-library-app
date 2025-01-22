package mock

import (
	"simple-library-app/module/library/entity"

	"github.com/stretchr/testify/mock"
)

type MockBookRepo struct {
	mock.Mock
}

func (m *MockBookRepo) GetBySubject(subject string) ([]*entity.Book, error) {
	args := m.Called(subject)

	if books, ok := args.Get(0).([]*entity.Book); ok {
		return books, args.Error(1)
	}

	return []*entity.Book{}, args.Error(1)
}

func (m *MockBookRepo) GetByEditionNumber(editionNumber string) (*entity.Book, error) {
	args := m.Called(editionNumber)
	if book, ok := args.Get(0).(*entity.Book); ok {
		return book, args.Error(1)
	}

	return nil, args.Error(1)
}
