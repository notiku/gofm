package gofm

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	Version = "0.0.1"
)

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

// Creates a new Client with the given API key.
func New(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}
}

const baseURL = "http://ws.audioscrobbler.com/2.0/"

// Handles the HTTP requests to the Last.fm API.
func (c *Client) doRequest(method string, params map[string]string, result interface{}) error {
	params["api_key"] = c.APIKey
	params["format"] = "json"

	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}
