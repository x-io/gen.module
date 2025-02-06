package dbs

import (
	"database/sql/driver"
)

//UUID string
type UUID string

//Value Value
func (j UUID) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "00000000-0000-0000-0000-000000000000", nil
	}
	return string(j), nil
}

//Scan Scan
func (j *UUID) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}

	switch v := src.(type) {
	case string:
		*j = UUID(v)
	case []byte:
		*j = UUID(v)
	}

	return nil
}

// // MarshalJSON returns j as the JSON encoding of j.
// func (j String) MarshalJSON() ([]byte, error) {
// 	if j == nil {
// 		return []byte("null"), nil
// 	}
// 	return j, nil
// }

// // UnmarshalJSON sets *j to a copy of data.
// func (j *String) UnmarshalJSON(data []byte) error {
// 	if j == nil {
// 		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
// 	}
// 	*j = append((*j)[0:0], data...)
// 	return nil
// }
