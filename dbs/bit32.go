package dbs

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// Bit32 int
type Bit32 int

// Value Value
func (j Bit32) Value() (driver.Value, error) {
	return fmt.Sprintf("%032b", j), nil
}

// Scan Scan
func (j *Bit32) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}
	switch v := src.(type) {
	case []byte:
		*j = convBit32(string(v))
	case int64:
		*j = Bit32(v)
	case string:
		*j = convBit32(v)
	}

	return nil
}

func convBit32(str string) Bit32 {
	str = strings.ReplaceAll(str, "\"", "")

	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		i, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	return Bit32(i)
}

// // MarshalJSON returns j as the JSON encoding of j.
// func (J Bool) MarshalJSON() ([]byte, error) {
// 	if j == nil {
// 		return []byte("null"), nil
// 	}
// 	return j, nil
// }

// UnmarshalJSON sets *j to a copy of data.
func (j *Bit32) UnmarshalJSON(data []byte) error {
	if j == nil {
		return fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = convBit32(string(data))
	return nil
}

func (j Bit32) Check(position int) bool {
	return j>>position&0x01 == 1
}
