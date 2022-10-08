package youtube

import "net/url"

type Request struct {
	Params map[string]string
}

func NewRequest() Request {
	return Request{
		Params: map[string]string{},
	}
}

func (request Request) AddParam(key, value string) Request {
	request.Params[key] = value

	return request
}

func (request Request) GetAsQuery() string {
	query := url.Values{}

	for key, value := range request.Params {
		query.Add(key, value)
	}

	return query.Encode()
}
