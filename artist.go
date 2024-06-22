package gofm

import (
	"regexp"
)

type Artist struct {
	Name  string `json:"name"`
	Mbid  string `json:"mbid"`
	URL   string `json:"url"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Streamable string `json:"streamable"`
	OnTour     string `json:"ontour"`
	Stats      struct {
		Listeners     string `json:"listeners"`
		Playcount     string `json:"playcount"`
		UserPlaycount string `json:"userplaycount"`
	} `json:"stats"`
	Similar struct {
		Artist []struct {
			Name  string `json:"name"`
			URL   string `json:"url"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
		} `json:"artist"`
	} `json:"similar"`
	Tags struct {
		Tag []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"tag"`
	} `json:"tags"`
	Bio struct {
		Published string `json:"published"`
		Summary   string `json:"summary"`
		Content   string `json:"content"`
	} `json:"bio"`
}

// Retrieves information about an artist.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
// Provide a username to also get the user's play count for the artist.
func (c *Client) GetArtistInfo(artist string, username ...string) (*Artist, error) {
	params := map[string]string{
		"method": "artist.getinfo",
	}

	if isValidMBID(artist) {
		params["mbid"] = artist
	} else {
		params["artist"] = artist
	}

	if len(username) > 0 {
		params["username"] = username[0]
	}

	var result struct {
		Artist Artist `json:"artist"`
	}

	err := c.doRequest("GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.Artist, nil
}

// Checks if the given string is a valid MBID (UUID format)
func isValidMBID(identifier string) bool {
	matched, _ := regexp.MatchString("^[a-f0-9-]{36}$", identifier)
	return matched
}
