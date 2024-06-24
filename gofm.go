package gofm

import (
	"net/http"

	ar "github.com/notiku/gofm/artist"
	us "github.com/notiku/gofm/user"
)

const (
	VERSION = "0.1.0" // Semantic versioning
	baseURL = "http://ws.audioscrobbler.com/2.0/"
	authURL = "http://www.last.fm/api/auth/"
)

type Network struct {
	APIKey     string
	APISecret  string
	BaseURL    string
	AuthURL    string
	HTTPClient *http.Client
	Artist     *ar.ArtistService
	User       *us.UserService
}

// Creates a new Network with the given API credentials.
func New(apiKey string, apiSecret string) *Network {
	return &Network{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		BaseURL:    baseURL,
		AuthURL:    authURL,
		HTTPClient: http.DefaultClient,
		Artist:     &ar.ArtistService{APIKey: apiKey, APISecret: apiSecret},
		User:       &us.UserService{APIKey: apiKey, APISecret: apiSecret},
	}
}

// Returns the version of the library.
func (c *Network) GetVersion() string {
	return VERSION
}
