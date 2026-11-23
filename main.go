package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

var videoExtensions = map[string]struct{}{
	"mp4": {},
	"mkv": {},
	"avi": {},
	"mov": {},
	"flv": {},
}

func main() {
	fmt.Println("mediahub starting up")
	files := []string{"movie.mp4", "documentary.mkv", "song.mp3", "clip.avi", "document.txt"}
	for _, file := range files {
		ext := filepath.Ext(file)
		if isVideo(ext) {
			fmt.Printf("Video file: %s\n", file)
		}
	}

	movie := describe("Serial Mom", 1994)
	fmt.Printf("Movie: %s\n", movie)

	year := "2001"
	yearInt, err := parseYear(year)
	if err != nil {
		fmt.Printf("Error parsing year: %v\n", err)
		return
	}
	fmt.Printf("Year: %d\n", yearInt)
}

func describe(title string, year int) string {
	return fmt.Sprintf("%s (%d)", title, year)
}

func isVideo(ext string) bool {

	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	_, ok := videoExtensions[ext]
	return ok
}

func parseYear(s string) (int, error) {
	var year int
	year, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid year format: %w", err)
	}
	if year < 1888 {
		return 0, fmt.Errorf("year %d is before the first known film", year)
	}
	return year, nil
}
