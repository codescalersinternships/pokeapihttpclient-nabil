package client

import (
	"net/http"
	"time"
)

var (
	DefaultApiUrl = "https://pokeapi.co"
)

type Option func(c *Client)
type Client struct {
	apiUrl     string
	httpClient *http.Client
}

func NewClient(timeout time.Duration, opts ...Option) *Client {
	client := &Client{
		apiUrl: DefaultApiUrl,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
	for _, o := range opts {
		o(client)
	}
	return client
}

func WithApiUrl(url string) Option {
	return func(c *Client) {
		c.apiUrl = url
	}
}
