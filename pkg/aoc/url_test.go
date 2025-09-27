package aoc_test

import (
	"lorech/advent-of-code/pkg/aoc"
	"testing"
)

func TestPuzzleUrl(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		day     int
		want    string
		wantErr bool
	}{
		{
			name:    "formats correctly",
			year:    2024,
			day:     25,
			want:    "https://adventofcode.com/2024/day/25",
			wantErr: false,
		},
		{
			name:    "single day has no leading zero",
			year:    2024,
			day:     1,
			want:    "https://adventofcode.com/2024/day/1",
			wantErr: false,
		},
		{
			name:    "errors when year before aoc start",
			year:    2010,
			day:     1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when year that hasn't happened",
			year:    3000,
			day:     1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when day invalid",
			year:    2024,
			day:     -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when day after christmas",
			year:    2024,
			day:     31,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := aoc.PuzzleUrl(tt.year, tt.day)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("PuzzleUrl() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("PuzzleUrl() succeeded unexpectedly")
			}
			if got == tt.want {
				t.Errorf("PuzzleUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputUrl(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		day     int
		want    string
		wantErr bool
	}{
		{
			name:    "formats correctly",
			year:    2024,
			day:     25,
			want:    "https://adventofcode.com/2024/day/25/input",
			wantErr: false,
		},
		{
			name:    "single day has no leading zero",
			year:    2024,
			day:     1,
			want:    "https://adventofcode.com/2024/day/1/input",
			wantErr: false,
		},
		{
			name:    "errors when year before aoc start",
			year:    2010,
			day:     1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when year that hasn't happened",
			year:    3000,
			day:     1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when day invalid",
			year:    2024,
			day:     -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "errors when day after christmas",
			year:    2024,
			day:     31,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := aoc.InputUrl(tt.year, tt.day)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("InputUrl() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("InputUrl() succeeded unexpectedly")
			}
			if got == tt.want {
				t.Errorf("InputUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
