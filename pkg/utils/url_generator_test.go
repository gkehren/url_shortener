package utils

import "testing"

func TestGeneratorShortURL(t *testing.T) {
	shortURL := GenerateShortURL()
	if len(shortURL) != 6 {
		t.Errorf("Expected short URL length to be 6, but got %d", len(shortURL))
	}
}
