package dbs

import (
	"database/sql/driver"
	"strings"
)

//String string
type String string

//Value Value
func (j String) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

//Scan Scan
func (j *String) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}

	switch v := src.(type) {
	case string:
		*j = String(v)
	case []byte:
		*j = String(v)
	}

	return nil
}

func (j String) String() string {
	return string(j)
}

func (j String) ToUpper() String {
	return String(strings.ToUpper(string(j)))
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
