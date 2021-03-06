// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Date allows for unmarshalling the date objects returned by Exact.
type Date struct {
	time.Time
}

// NewDate generates a new Date.
func NewDate(t *time.Time) Date {
	d := Date{}
	if t != nil {
		d.Time = *t
	}
	return d
}

// NewDatePtr generates a pointer to a new Date.
func NewDatePtr(t *time.Time) *Date {
	d := NewDate(t)
	return &d
}

// IsSet returns a boolean if the Date is actually set.
func (d *Date) IsSet() bool {
	return !d.IsZero()
}

// UnmarshalJSON unmarshals the date format returned from the
// Exact Online API.
func (d *Date) UnmarshalJSON(b []byte) error {
	re := regexp.MustCompile(`[0-9]+`)
	s := re.FindString(string(b))
	if s == "" {
		d.Time = time.Time{}
		return nil
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("Date.UnmarshalJSON() error: %v", err)
	}

	d.Time = time.Unix(0, i*int64(time.Millisecond))
	return nil
}

// MarshalJSON marshals the date to a format expected by the
// Exact Online API.
func (d *Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}

	return d.Time.MarshalJSON()
}

// StartDate returns passed date of '1800-01-01' if passed date is nil
func (d *Date) StartDate() Date {
	if d != nil {
		return *d
	}

	startDate, _ := time.Parse("2006-01-02", "1800-01-01")
	return Date{startDate}
}

// EndDate returns passed date of '2099-12-31' if passed date is nil
func (d *Date) EndDate() Date {
	if d != nil {
		return *d
	}

	endDate, _ := time.Parse("2006-01-02", "2099-12-31")
	return Date{endDate}
}

func (d1 Date) Before(d2 Date) bool {
	return d1.Time.Before(d2.Time)
}

func (d1 Date) After(d2 Date) bool {
	return d1.Time.After(d2.Time)
}

func (d1 Date) Between(d2 Date, d3 Date) bool {
	return (d1.Time.After(d2.Time) && d1.Time.Before(d3.Time)) || (d1.Time.After(d3.Time) && d1.Time.Before(d2.Time))
}
