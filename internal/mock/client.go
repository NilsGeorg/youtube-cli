package mock

import (
	"bytes"
	"io"
	"net/http"
)

type HTTPClient struct {
	JSON string
}

func (httpClient HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{Body: io.NopCloser(bytes.NewBufferString(httpClient.JSON))}, nil
}
