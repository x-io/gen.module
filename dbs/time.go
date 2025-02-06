package dbs

import (
	"database/sql/driver"
	"time"
)

// Time int
type Time int64

// Value Value
func (j Time) Value() (driver.Value, error) {
	if j == 0 {
		return nil, nil
	}
	return time.Unix(int64(j), 0), nil
}

// Scan Scan
func (j *Time) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}

	switch v := src.(type) {
	case int:
		*j = Time(v)
	case int64:
		*j = Time(v)
	case float64:
		*j = Time(v)
	case time.Time:
		*j = Time(v.Unix())
	}
	return nil
}

// Time Time
func (j Time) Time() time.Time {
	if j == 0 {
		return time.Unix(0, 0)
	}

	return time.Unix(int64(j), 0)
}

// Time Time
func (j *Time) Date() (year int, month time.Month, day int) {
	if j == nil {
		return 0, 0, 0
	}

	return time.Unix(int64(*j), 0).Date()
}
