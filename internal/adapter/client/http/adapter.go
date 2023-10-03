package httpclient

import (
	"errors"
	"net/http"
)

type HttpAdapter struct {
	endpoint string
	client   *http.Client
}

var (
	ErrorBillNotFound = errors.New("cannot find bill number")
)

func NewHttpAdapter(endpoint string) *HttpAdapter {

	httpClient := &http.Client{}

	return &HttpAdapter{
		endpoint: endpoint,
		client:   httpClient,
	}
}
