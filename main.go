package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type Movie struct {
	Title     string
	Year      int
	Path      string
	Monitored bool
}

var videoExtensions = map[string]struct{}{
	"mp4": {},
	"mkv": {},
	"avi": {},
	"mov": {},
	"flv": {},
}

func main() {
	fmt.Println("mediahub starting up")
	movie := Movie{
		Title:     "Serial Mom",
		Year:      1994,
		Path:      "/path/to/serial_mom.mp4",
		Monitored: false,
	}
	movies := []*Movie{
		{Title: "Inception", Year: 2010, Path: "/path/to/inception.mp4", Monitored: true},
		{Title: "The Matrix", Year: 1999, Path: "/path/to/the_matrix.mkv", Monitored: false},
		{Title: "Interstellar", Year: 2014, Path: "/path/to/interstellar.avi", Monitored: true},
	}
	monitoredMovies := monitored(movies)

	for _, m := range monitoredMovies {
		fmt.Printf("%s (%d)\n", m.Title, m.Year)
	}
	fmt.Printf("Movie: %s\n", movie.Describe())

	files := []string{"movie.mp4", "documentary.mkv", "song.mp3", "clip.avi", "document.txt"}
	for _, file := range files {
		ext := filepath.Ext(file)
		if isVideo(ext) {
			fmt.Printf("Video file: %s\n", file)
		}
	}

	year := "2001"
	yearInt, err := parseYear(year)
	if err != nil {
		fmt.Printf("Error parsing year: %v\n", err)
		return
	}
	fmt.Printf("Year: %d\n", yearInt)
}

func (m *Movie) Describe() string {
	return fmt.Sprintf("%s (%d)", m.Title, m.Year)
}

func (m *Movie) Monitor() {
	m.Monitored = true

}

func (m *Movie) Unmonitor() {
	m.Monitored = false

}

func NewMovie(title, yearStr, path string) (*Movie, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	year, err := parseYear(yearStr)
	if err != nil {
		return nil, fmt.Errorf("invalid year: %w", err)
	}
	if !isVideo(filepath.Ext(path)) {
		return nil, fmt.Errorf("file %q is not a recognized video format", path)
	}
	return &Movie{
		Title:     title,
		Year:      year,
		Path:      path,
		Monitored: false,
	}, nil
}
func isVideo(ext string) bool {

	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	_, ok := videoExtensions[ext]
	return ok
}

func parseYear(s string) (int, error) {
	year, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid year format: %w", err)
	}
	if year < 1888 {
		return 0, fmt.Errorf("year %d is before the first known film", year)
	}
	return year, nil
}

func monitored(movies []*Movie) []*Movie {
	var result []*Movie
	for _, m := range movies {
		if m != nil && m.Monitored {
			result = append(result, m)
		}
	}
	return result
}
