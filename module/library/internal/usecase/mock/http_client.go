package mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type HTTPClient struct {
	mock.Mock
}

func (m *HTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
