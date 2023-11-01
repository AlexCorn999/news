package news

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const APIKEY = "88ea8ecf6571e60725e84639289b3ed1"

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetCategory(category Category) (*NewsResponse, error) {
	link := fmt.Sprintf("https://gnews.io/api/v4/top-headlines?category=%s&lang=ru&country=ru&max=5&apikey=%s", category, APIKEY)
	resp, err := c.client.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var n NewsResponse
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, err
	}
	return &n, nil
}

func (c *Client) GetNewsByCountry(category Category, country Country) (*NewsResponse, error) {
	link := fmt.Sprintf("https://gnews.io/api/v4/top-headlines?category=%s&lang=%s&country=ru&max=5&apikey=%s", category, country, APIKEY)
	resp, err := c.client.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var n NewsResponse
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, err
	}
	return &n, nil
}
