package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "github.com/mattn/go-mastodon"
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

func main() {
	filename := "./data/dailyintegral.json"
	puzzles := parse_puzzles(filename)

	for _, puzzle := range puzzles.Content {
		fmt.Println("diff:", puzzle.Difficulty)
	}
}
