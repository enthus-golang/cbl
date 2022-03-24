package cbl

import (
	"errors"
	"time"
)

const (
	DateFormat     = "20060102"
	DateTimeFormat = "20060102T15:04:05"
)

type Date struct {
	Time time.Time
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (d Date) MarshalText() ([]byte, error) {
	if y := d.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateFormat))
	return d.Time.AppendFormat(b, DateFormat), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (d *Date) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*d).Time, err = time.Parse(DateFormat, string(data))
	return err
}

type DateTime struct {
	Time time.Time
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (d DateTime) MarshalText() ([]byte, error) {
	if y := d.Time.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateTimeFormat))
	return d.Time.AppendFormat(b, DateTimeFormat), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (d *DateTime) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*d).Time, err = time.Parse(DateTimeFormat, string(data))
	return err
}
