package dbs

import (
	"database/sql/driver"
)

//Bool db Type
type Bool bool

//Value Value
func (j Bool) Value() (driver.Value, error) {
	return bool(j), nil
}

//Scan Scan
func (j *Bool) Scan(src interface{}) error {
	if src == nil {
		*j = false
		return nil
	}

	switch v := src.(type) {
	case bool:
		*j = Bool(v)
	case string:
		*j = Bool(v == "true")
	}

	return nil
}

// // MarshalJSON returns j as the JSON encoding of j.
// func (J Bool) MarshalJSON() ([]byte, error) {
// 	if j == nil {
// 		return []byte("null"), nil
// 	}
// 	return j, nil
// }

// // UnmarshalJSON sets *j to a copy of data.
// func (j *Int) UnmarshalJSON(data []byte) error {
// 	if j == nil {
// 		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
// 	}
// 	*j = append((*j)[0:0], data...)
// 	return nil
// }
