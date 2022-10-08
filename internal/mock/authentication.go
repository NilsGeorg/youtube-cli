package mock

var AuthenticationToken = "MOCK_TOKEN"

type Authentication struct{}

func (m Authentication) GetToken() (string, error) {
	return AuthenticationToken, nil
}

func (m Authentication) SetToken(token string) error {
	return nil
}
