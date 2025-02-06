package dbs

import (
	"database/sql/driver"
)

//String string
type Strings []string

//Value Value
func (j Strings) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return []string(j), nil
}

//Scan Scan
func (j *Strings) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case []string:
		*j = Strings(v)
		// case []byte:
		// 	*j = Strings(v)
	}

	return nil
}

func (j Strings) String() []string {
	return []string(j)
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
