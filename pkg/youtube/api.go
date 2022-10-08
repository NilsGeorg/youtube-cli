package youtube

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/NilsGeorg/ycli/pkg/config"
)

type API[T ResponseType] interface {
	Get(path string, request Request) (*Response[T], error)
}

type APIImpl[T ResponseType] struct {
	HTTPClient     HTTPClient
	Authentication config.Authentication
	YoutubeAPIURL  string
	YoutubeAPIPath string
}

func NewAPI[T ResponseType](c Client) APIImpl[T] {
	api := APIImpl[T]{}
	api.Authentication = c.Authentication
	api.HTTPClient = c.HTTPClient
	api.YoutubeAPIPath = c.YoutubeAPIPath
	api.YoutubeAPIURL = c.YoutubeAPIURL

	return api
}

func (api APIImpl[T]) Get(path string, request Request) (*Response[T], error) {
	httpRequest, err := api.createRequest(path, request)
	if err != nil {
		return nil, err
	}

	response, err := api.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return api.responseToStruct(response)
}

func (api APIImpl[T]) createRequest(path string, request Request) (*http.Request, error) {
	if path[0] != '/' {
		path = "/" + path
	}

	token, err := api.Authentication.GetToken()
	if err != nil {
		return nil, err
	}

	request.AddParam("key", token)

	req, err := http.NewRequest("GET", "https://"+api.YoutubeAPIURL+api.YoutubeAPIPath+path, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = request.GetAsQuery()

	return req, nil
}

func (api APIImpl[T]) responseToStruct(httpResponse *http.Response) (*Response[T], error) {
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var response *Response[T]
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	if response.Items == nil {
		response.Items = []*T{}
	}

	return response, nil
}
