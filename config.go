package circleci

import (
	"errors"
	"net/http"
)

func NewHTTPClient(token string) *http.Client {
	c := &Config{
		APIToken: token,
	}

	return c.Client()
}

type Config struct {
	APIToken string
}

func (c *Config) Client() *http.Client {
	return &http.Client{
		Transport: &Transport{
			Config: c,
		},
	}
}

func (c *Config) Token() (string, error) {
	if c.APIToken == "" {
		return "", errors.New("circleci: Not found Token.")
	}

	return c.APIToken, nil
}
