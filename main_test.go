package main

import "testing"

func TestYearParsing(t *testing.T) {
	yearTests := []struct {
		input   string
		wantErr bool
	}{
		{"2021", false},
		{"abc", true},
		{"1850", true},
		{"", true},
	}

	for _, tt := range yearTests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parseYear(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYear(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestNewMovie(t *testing.T) {
	movieTests := []struct {
		title string
		year  string
		path  string
	}{
		{"", "2010", "inception.mkv"},
		{"South Park", "nineteen ninety-nine", "south_park.mp4"},
		{"The Matrix", "1999", "the_matrix.txt"},
		{"Interstellar", "2014", "interstellar.avi"},
	}

	for _, tt := range movieTests {
		t.Run(tt.title, func(t *testing.T) {
			_, err := NewMovie(tt.title, tt.year, tt.path)
			if err != nil {
				t.Errorf("newMovie(%q, %q, %q) error = %v", tt.title, tt.year, tt.path, err)
			}
		})
	}
}
func TestIsVideo(t *testing.T) {
	videoTests := []struct {
		input string
		want  bool
	}{
		{".mkv", true},
		{".MKV", true},
		{"mkv", true},
		{".mp3", false},
		{"", false},
	}

	for _, tt := range videoTests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isVideo(tt.input); got != tt.want {
				t.Errorf("isVideo(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMonitor(t *testing.T) {
	movie := &Movie{Title: "Inception", Year: 2010, Path: "/path/to/inception.mp4", Monitored: false}
	movie.Monitor()
	if !movie.Monitored {
		t.Errorf("Monitor() failed, expected Monitored to be true")
	}
}
