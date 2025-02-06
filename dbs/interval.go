package dbs

import (
	"bytes"
	"database/sql/driver"
	"log"
)

//Interval string
type Interval string

//Value Value
func (j Interval) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return string(j), nil
}

//Scan Scan
func (j *Interval) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}

	switch v := src.(type) {
	case []byte:
		if bytes.Equal(v, []byte{48, 48, 58, 48, 48, 58, 48, 48}) {
			*j = ""
		} else {
			*j = Interval(v)
		}
	case string:
		if v == "00:00:00" {
			*j = ""
		} else {
			*j = Interval(v)
		}
	default:
		log.Println(src)
	}

	return nil
}
