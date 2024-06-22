package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/notiku/gofm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	apiKey := os.Getenv("LASTFM_API_KEY")
	c := gofm.New(apiKey)

	// Get Artist info
	artist, err := c.GetArtistInfo("Cher", "RJ")
	if err != nil {
		log.Fatalf("Error getting artist info: %v", err)
	}

	fmt.Printf("Name: %s\nURL: %s\nPlaycount: %s\nListeners: %s\n", artist.Name, artist.URL, artist.Stats.Playcount, artist.Stats.Listeners)

	// UserPlaycount is only available if a username is provided in GetArtistInfo
	fmt.Printf("User Playcount: %s\n", artist.Stats.UserPlaycount)
}
