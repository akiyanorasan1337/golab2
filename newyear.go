// Package main содержит приложение для расчёта количества дней
// до ближайшего следующего Нового года.
package main

import (
	"errors"
	"time"
)

// ErrNilDate возвращается, если передана нулевая дата.
var ErrNilDate = errors.New("date is zero")

// DaysUntilNewYear возвращает количество календарных дней,
// оставшихся от даты d до ближайшего следующего 1 января.
func DaysUntilNewYear(d time.Time) (int, error) {
	if d.IsZero() {
		return 0, ErrNilDate
	}

	d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())

	if d.Month() == time.January && d.Day() == 1 {
		return 0, nil
	}

	next := time.Date(d.Year()+1, time.January, 1, 0, 0, 0, 0, d.Location())
	days := int(next.Sub(d).Hours() / 24)
	return days, nil
}
