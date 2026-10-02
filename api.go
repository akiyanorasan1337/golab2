package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const dateLayout = "2006-01-02"

// Response — машиночитаемый ответ API.
type Response struct {
	Date  string `json:"date"`
	Days  int    `json:"days"`
	Error string `json:"error,omitempty"`
}

// NewYearHandler обрабатывает запросы на получение количества дней до Нового года.
func NewYearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "method not allowed"})
		return
	}

	dateStr := r.URL.Query().Get("date")
	var date time.Time
	if dateStr == "" {
		date = time.Now()
	} else {
		parsed, err := time.Parse(dateLayout, dateStr)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, Response{Error: "invalid date format, expected YYYY-MM-DD"})
			return
		}
		date = parsed
	}

	days, err := DaysUntilNewYear(date)
	if err != nil {
		if errors.Is(err, ErrNilDate) {
			writeJSON(w, http.StatusBadRequest, Response{Error: "invalid date"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, Response{Error: "internal error"})
		return
	}

	writeJSON(w, http.StatusOK, Response{Date: date.Format(dateLayout), Days: days})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
