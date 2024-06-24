package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

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
	apiSecret := os.Getenv("LASTFM_API_SECRET")
	// Create a new Network
	n := gofm.New(apiKey, apiSecret)

	fmt.Println("Version: ", n.GetVersion())

	// Get Artist info
	artist, err := n.Artist.GetInfo("1nonly", "unb_")
	if err != nil {
		log.Fatalf("Error getting artist info: %v", err)
	}

	fmt.Printf("Playcount: %s, Listeners: %s\n", artist.Stats.Playcount, artist.Stats.Listeners)

	// UserPlaycount is only available if a username is provided
	fmt.Printf("User Playcount: %s\n", artist.Stats.UserPlaycount)

	// Get similar artists
	similar, err := n.Artist.GetSimilar("1nonly")
	if err != nil {
		log.Fatalf("Error getting similar artists: %v", err)
	}

	// Randomly select a similar artist
	ranIndex := 0
	if len(similar.Artist) > 1 {
		ranIndex = rand.Intn(len(similar.Artist))
	}

	match := similar.Artist[ranIndex].Match
	matchInt, err := strconv.ParseFloat(match, 64)
	if err != nil {
		log.Fatalf("Error parsing match: %v", err)
	}

	matchPercent := matchInt * 100
	fmt.Printf("Similar artist: %s [Match: %.2f%%]\n", similar.Artist[ranIndex].Name, matchPercent)
}
