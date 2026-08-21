package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

type Puzzles struct {
	Content []Puzzle `json:"puzzles"`
}

type Puzzle struct {
	Day        int    `json:"day"`
	Difficulty string `json:"difficulty"`
	Title      string `json:"title"`
	Latex      string `json:"latex"`
}

func parse_puzzles(filename string) Puzzles {
	byte_data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading the file: %v\n", err)
		return Puzzles{}
	}

	var puzzles Puzzles
	err = json.Unmarshal(byte_data, &puzzles)
	if err != nil {
		fmt.Printf("Error unmarshalling json: %v\n", err)
		return Puzzles{}
	}

	return puzzles
}

func get_config_from_env() mastodon.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := mastodon.Config{
		Server: "https://mathstodon.xyz",
	}

	cfg.ClientID = os.Getenv("CLIENT_KEY")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")
	cfg.AccessToken = os.Getenv("ACCESS_TOKEN")

	return cfg
}

func create_toot(p Puzzle) string {
	integral_latex, _, _ := strings.Cut(p.Latex, "\n")
	diff := strings.ToLower(p.Difficulty)
	toot := fmt.Sprintf("%s (day no. %d)\nDifficulty: %s\n\n\\(%s\\)\n\nHave fun!", p.Title, p.Day, diff, integral_latex)

	return toot
}

func main() {
	filename := "./data/dailyintegral.json"
	puzzles := parse_puzzles(filename)

	cfg := get_config_from_env()
	c := mastodon.NewClient(&cfg)

	for i := 0; i < len(puzzles.Content); i++ {
		_, err := c.PostStatus(context.Background(), &mastodon.Toot{
			Status: create_toot(puzzles.Content[i]),
		})

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(5 * time.Second)
		log.Printf("toot posted! (%d)\n", i+1)
	}
}
