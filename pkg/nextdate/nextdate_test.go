package nextdate

import (
	"testing"
	"time"
)

func TestNextDate(t *testing.T) {
	now := time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		start    string
		repeat   string
		expected string
	}{
		{
			name:     "every 7 days",
			start:    "20240113",
			repeat:   "d 7",
			expected: "20240127",
		},
		{
			name:     "every year",
			start:    "20240229",
			repeat:   "y",
			expected: "20250301",
		},
		{
			name:     "every day",
			start:    "20240125",
			repeat:   "d 1",
			expected: "20240127",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextDate(now, tt.start, tt.repeat)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}
