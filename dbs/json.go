package dbs

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
)

//JSON Json Raw
type JSON []byte

//Value Value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return []byte(j), nil
}

//Scan Scan
func (j *JSON) Scan(src interface{}) error {
	if src == nil {
		*j = []byte{}
		return nil
	}

	switch v := src.(type) {
	case string:
		*j = []byte(fmt.Sprintf("\"%s\"", v))
	case []byte:
		if bytes.Equal(v, []byte{123, 125}) || bytes.Equal(v, []byte{91, 93}) {
			*j = []byte{}
			return nil
		}
		*j = append((*j)[0:0], v...)
	}

	return nil
}

// MarshalJSON returns j as the JSON encoding of j.
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON sets *j to a copy of data.
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
