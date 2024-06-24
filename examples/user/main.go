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
	apiSecret := os.Getenv("LASTFM_API_SECRET")
	// Create a new Network
	n := gofm.New(apiKey, apiSecret)

	fmt.Println("Version: ", n.GetVersion())

	// Get user info
	user, err := n.User.GetInfo("unb_")
	if err != nil {
		log.Fatal("There was an error getting user info: ", err)
	}

	fmt.Printf("User: %s\nPlaycount: %s\n", user.Name, user.Playcount)
}
