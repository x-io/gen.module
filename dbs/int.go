package dbs

import (
	"database/sql/driver"
	"strconv"
)

// Int int
type Int int64

// Value Value
func (j Int) Value() (driver.Value, error) {
	return int64(j), nil
}

// Scan Scan
func (j *Int) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}
	switch v := src.(type) {
	case int:
		*j = Int(v)
	case int64:
		*j = Int(v)
	case float32:
		*j = Int(v)
	case float64:
		*j = Int(v)
	case string:
		vv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*j = Int(vv)
	}

	return nil
}
