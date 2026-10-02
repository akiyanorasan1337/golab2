package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewYearHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantErr    bool
	}{
		{"валидная дата", "/days?date=2024-12-31", http.StatusOK, false},
		{"1 января", "/days?date=2024-01-01", http.StatusOK, false},
		{"некорректный формат", "/days?date=31-12-2024", http.StatusBadRequest, true},
		{"без даты", "/days", http.StatusOK, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rec := httptest.NewRecorder()
			NewYearHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantStatus {
				t.Fatalf("status = %d, want %d", res.StatusCode, tt.wantStatus)
			}

			var resp Response
			if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
				t.Fatalf("decode error: %v", err)
			}

			if tt.wantErr && resp.Error == "" {
				t.Errorf("expected error message")
			}
		})
	}
}
