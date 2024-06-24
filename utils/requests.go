package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "http://ws.audioscrobbler.com/2.0/"

// Handles the HTTP requests to the Last.fm API.
func DoRequest(apiKey, method string, params map[string]string, result interface{}) error {
	params["api_key"] = apiKey
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

	client := &http.Client{}
	resp, err := client.Do(req)
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

// Handles the hashed HTTP requests to the Last.fm API.
func DoHashedRequest(apiKey, apiSecret, method string, params map[string]string, result interface{}) error {
	hash := HashRequest(params, apiSecret)
	params["api_sig"] = hash
	params["api_key"] = apiKey
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

	client := &http.Client{}
	resp, err := client.Do(req)
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
