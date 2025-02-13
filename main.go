package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type URLMapping struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

var mappings []URLMapping

func loadMappings() {
	file, err := os.Open("urls.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&mappings); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
}

func saveMappings() {
	file, err := os.Create("urls.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(mappings); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seeedRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seeedRand.Intn(len(charset))]
	}
	return string(b)
}

func shortenURL(longURL string) string {
	if shortURL, found := resolveLongURL(longURL); found {
		return shortURL
	} else {
		shortURL := generateShortURL()
		mappings = append(mappings, URLMapping{ShortURL: shortURL, LongURL: longURL})
		saveMappings()
		return shortURL
	}
}

func resolveURL(shortURL string) (string, bool) {
	for _, mapping := range mappings {
		if mapping.ShortURL == shortURL {
			return mapping.LongURL, true
		}
	}
	return "", false
}

func resolveLongURL(longURL string) (string, bool) {
	for _, mapping := range mappings {
		if mapping.LongURL == longURL {
			return mapping.ShortURL, true
		}
	}
	return "", false
}

func main() {
	loadMappings()

	var command, url string
	fmt.Println("Enter command (shorten/resolve/exit):")
	for {
		fmt.Scanln(&command)
		switch command {
		case "shorten":
			fmt.Println("Enter URL to shorten:")
			fmt.Scanln(&url)
			shortURL := shortenURL(url)
			fmt.Printf("Shortened URL: %s\n", shortURL)
		case "resolve":
			fmt.Println("Enter a short URL to resolve:")
			fmt.Scanln(&url)
			if longURL, found := resolveURL(url); found {
				fmt.Printf("Resolved URL: %s\n", longURL)
			} else {
				fmt.Println("Short URL not found.")
			}
		case "exit":
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
