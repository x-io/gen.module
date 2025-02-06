package dbs

import (
	"database/sql/driver"
)

//Point string
type Point string

//Value Value
func (j Point) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

//Scan Scan
func (j *Point) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}

	switch v := src.(type) {
	case []byte:
		*j = Point(v[1 : len(v)-1])
	case string:
		*j = Point(v[1 : len(v)-1])		
	}

	return nil
}
