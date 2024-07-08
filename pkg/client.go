package client

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
	return client
}
