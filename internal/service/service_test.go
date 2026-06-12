package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		now      time.Time
		expected int
	}{
		{
			name:     "Birthday is today",
			dob:      time.Date(1990, 6, 11, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2026, 6, 11, 12, 0, 0, 0, time.UTC),
			expected: 36,
		},
		{
			name:     "Birthday was yesterday",
			dob:      time.Date(1990, 6, 10, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2026, 6, 11, 12, 0, 0, 0, time.UTC),
			expected: 36,
		},
		{
			name:     "Birthday is tomorrow",
			dob:      time.Date(1990, 6, 12, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2026, 6, 11, 12, 0, 0, 0, time.UTC),
			expected: 35,
		},
		{
			name:     "Future DOB returns 0",
			dob:      time.Date(2028, 6, 11, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2026, 6, 11, 12, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "Leap year - born Feb 29, now is Feb 28 common year",
			dob:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2001, 2, 28, 12, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "Leap year - born Feb 29, now is March 1 common year",
			dob:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			now:      time.Date(2001, 3, 1, 12, 0, 0, 0, time.UTC),
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalculateAge(tt.dob, tt.now)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d for dob %v and now %v", tt.expected, actual, tt.dob, tt.now)
			}
		})
	}
}
