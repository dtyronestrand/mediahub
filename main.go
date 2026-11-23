package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("mediahub starting up")
	files := []string{"movie.mp4", "documentary.mkv", "song.mp3", "clip.avi", "document.txt"}
	for _, file := range files {
		ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file), "."))
		if isVideo(ext) {
			fmt.Printf("Video file: %s\n", file)
		}
	}
}

func describe(title string, year int) string {
	return fmt.Sprintf("%s (%d)", title, year)
}

func isVideo(ext string) bool {
	videoExtensions := []string{"mp4", "mkv", "avi", "mov", "flv"}
	return slices.Contains(videoExtensions, ext)
}

func parseYear(s string) (int, error) {
	var year int
	year, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid year format: %v", err)
	}
	return year, nil
}
