package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-library-app/module/library/entity"
)

type subjectBook struct {
	Title   string `json:"title"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Availability struct {
		ISBN string `json:"isbn"`
	} `json:"availability"`
}

type subjectApiResponse struct {
	Name  string        `json:"name"`
	Works []subjectBook `json:"works"`
}

func (sar *subjectApiResponse) toEntities() []*entity.Book {
	var books []*entity.Book

	for _, work := range sar.Works {
		authors := make([]string, len(work.Authors))
		for i, author := range work.Authors {
			authors[i] = author.Name
		}

		bookEntity := &entity.Book{
			Title:         work.Title,
			Authors:       authors,
			EditionNumber: work.Availability.ISBN,
		}

		books = append(books, bookEntity)
	}

	return books
}

type searchBook struct {
	Title   string   `json:"title"`
	Authors []string `json:"author_name"`
}

type searchApiResponse struct {
	NumFound int          `json:"numFound"`
	Docs     []searchBook `json:"docs"`
}

func (sar *searchApiResponse) toEntity() *entity.Book {
	if sar.NumFound > 0 {
		book := sar.Docs[0]
		return &entity.Book{
			Title:   book.Title,
			Authors: book.Authors,
		}
	}

	return nil
}

type BookRepository struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewBookRepository(baseUrl string, client *http.Client) *BookRepository {
	return &BookRepository{
		BaseURL:    baseUrl,
		HTTPClient: client,
	}
}

func (r *BookRepository) GetBySubject(subject string) ([]*entity.Book, error) {
	url := fmt.Sprintf("%s/%s.json", r.BaseURL, subject)

	resp, err := r.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp subjectApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	books := apiResp.toEntities()
	return books, nil
}

func (r *BookRepository) GetByEditionNumber(editionNumber string) (*entity.Book, error) {
	url := fmt.Sprintf("%s/search.json?isbn=%s", r.BaseURL, editionNumber)

	resp, err := r.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch book: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp searchApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	book := apiResp.toEntity()
	if book != nil {
		book.EditionNumber = editionNumber
	}

	return book, nil
}
