package artist

import (
	"github.com/notiku/gofm/utils"
)

// Retrieves information about an artist.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
// Provide a username to also get the user's play count for the artist.
func (c *ArtistService) GetInfo(artist string, username ...string) (*Artist, error) {
	params := map[string]string{
		"method": "artist.getinfo",
	}

	if utils.IsValidMBID(artist) {
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

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.Artist, nil
}

// Retrieves similar artists for a given artist.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
func (c *ArtistService) GetSimilar(artist string, limit ...string) (*Similar, error) {
	params := map[string]string{
		"method": "artist.getsimilar",
	}

	if len(limit) > 0 {
		params["limit"] = string(limit[0])
	}

	if utils.IsValidMBID(artist) {
		params["mbid"] = artist
	} else {
		params["artist"] = artist
	}

	var result struct {
		Similar Similar `json:"similarartists"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.Similar, nil
}

// Retrieves the tags applied by an individual user to an artist on Last.fm.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
func (c *ArtistService) GetTags(artist, user string) (*Tags, error) {
	params := map[string]string{
		"method": "artist.gettags",
		"artist": artist,
		"user":   user,
	}

	var result struct {
		Tags Tags `json:"tags"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.Tags, nil
}

// Retrieves the top tags for an artist on Last.fm, ordered by popularity.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
func (c *ArtistService) GetTopTags(artist string) (*Tags, error) {
	params := map[string]string{
		"method": "artist.gettoptags",
		"artist": artist,
	}

	var result struct {
		Tags Tags `json:"toptags"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.Tags, nil
}

// Retrieves the top tracks by an artist on Last.fm, ordered by popularity.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
func (c *ArtistService) GetTopTracks(artist string, limit ...string) (*Tracks, error) {
	params := map[string]string{
		"method": "artist.gettoptracks",
		"artist": artist,
	}

	if len(limit) > 0 {
		params["limit"] = limit[0]
	}

	var result struct {
		TopTracks Tracks `json:"toptracks"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.TopTracks, nil
}

// Retrieves the top albums by an artist on Last.fm, ordered by popularity.
// The artist can be identified by either its name or its MusicBrainz ID (MBID).
func (c *ArtistService) GetArtistTopAlbums(artist string, limit ...string) (*Albums, error) {
	params := map[string]string{
		"method": "artist.gettopalbums",
		"artist": artist,
	}

	if len(limit) > 0 {
		params["limit"] = limit[0]
	}

	var result struct {
		TopAlbums Albums `json:"topalbums"`
	}

	err := utils.DoRequest(c.APIKey, "GET", params, &result)
	if err != nil {
		return nil, err
	}

	return &result.TopAlbums, nil
}
