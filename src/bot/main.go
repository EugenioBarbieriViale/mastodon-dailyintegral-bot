package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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
	lines := strings.Split(p.Latex, "\n")
	var clean_latex []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			clean_latex = append(clean_latex, l)
		}
	}

	body := strings.Join(clean_latex, " \\\\\n")
	body = strings.ReplaceAll(body, `\[3pt]`, "")

	var latex strings.Builder
	latex.WriteString("\\[\n")
	latex.WriteString("\\begin{array}{c}\n")
	latex.WriteString(body)
	latex.WriteString("\n\\end{array}\n")
	latex.WriteString("\\]")

	diff := strings.ToLower(p.Difficulty)
	toot := fmt.Sprintf("%s (day no. %d)\nDifficulty: %s\n\n%s\n\nHave fun!", p.Title, p.Day, diff, latex.String())

	return toot
}

func main() {
	filename := "./data/dailyintegral.json"
	puzzles := parse_puzzles(filename)

	cfg := get_config_from_env()
	c := mastodon.NewClient(&cfg)

	_, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status: create_toot(puzzles.Content[0]),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("toot posted!")
}
