package circleci

import "net/http"

func NewClient(config *Config) *Client {
	return &Client{config.Client()}
}

type Client struct {
	hc *http.Client
}

func (c *Client) User() {
}
