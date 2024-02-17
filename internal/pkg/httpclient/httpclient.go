package httpclient

import (
	"encoding/json"
	"net/http"
)

type HttpClientInterface interface {
	Get(endpoint string, responseObj interface{}) *HttpClientError
}

type HttpClientError struct {
	Error      error
	StatusCode *int
}

type HttpClient struct{}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c HttpClient) Get(addr string, responseObj interface{}) *HttpClientError {
	req, err := http.NewRequest("GET", addr, nil)

	if err != nil {
		return &HttpClientError{
			Error: err,
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &HttpClientError{
			Error:      err,
			StatusCode: &resp.StatusCode,
		}
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return &HttpClientError{
			Error:      err,
			StatusCode: &resp.StatusCode,
		}
	}

	return nil
}
