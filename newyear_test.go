package main

import (
	"errors"
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	tests := []struct {
		name    string
		input   time.Time
		want    int
		wantErr error
	}{
		{"начало календарного года", time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), 0, nil},
		{"конец календарного года", time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC), 1, nil},
		{"високосный год — до 29 февраля", time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC), 308, nil},
		{"високосный год — после 29 февраля", time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC), 306, nil},
		{"невисокосный год — после 28 февраля", time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC), 306, nil},
		{"середина года", time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC), 184, nil},
		{"нулевая дата", time.Time{}, 0, ErrNilDate},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysUntilNewYear(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
