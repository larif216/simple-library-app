package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-library-app/module/library/internal/entity"
)

type book struct {
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Availability struct {
		ISBN              string `json:"isbn"`
		AvailableToBorrow bool   `json:"available_to_borrow"`
	} `json:"availability"`
}

type apiResponse struct {
	Name  string `json:"name"`
	Works []book `json:"works"`
}

func (ar *apiResponse) toEntities() []*entity.Book {
	var books []*entity.Book

	for _, work := range ar.Works {
		authors := make([]string, len(work.Authors))
		for i, author := range work.Authors {
			authors[i] = author.Name
		}

		bookEntity := &entity.Book{
			Title:         work.Title,
			Authors:       authors,
			EditionNumber: work.Availability.ISBN,
			IsAvailable:   work.Availability.AvailableToBorrow,
		}

		books = append(books, bookEntity)
	}

	return books
}

type BookRepository struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewBookRepository(baseUrl string, client *http.Client) BookRepository {
	return BookRepository{
		BaseURL:    baseUrl,
		HTTPClient: client,
	}
}

func (r BookRepository) GetBySubject(subject string) ([]*entity.Book, error) {
	url := fmt.Sprintf("%s/%s.json", r.BaseURL, subject)

	resp, err := r.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	books := apiResp.toEntities()
	return books, nil
}
