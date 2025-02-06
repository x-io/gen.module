package dbs

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// Bit16 int
type Bit16 int

// Value Value
func (j Bit16) Value() (driver.Value, error) {
	return fmt.Sprintf("%016b", j), nil
}

// Scan Scan
func (j *Bit16) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}
	switch v := src.(type) {
	case []byte:
		*j = convBit16(string(v))
	case int64:
		*j = Bit16(v)
	case string:
		*j = convBit16(v)
	}

	return nil
}

func convBit16(str string) Bit16 {
	str = strings.ReplaceAll(str, "\"", "")

	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		i, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	return Bit16(i)
}

// // MarshalJSON returns j as the JSON encoding of j.
// func (J Bool) MarshalJSON() ([]byte, error) {
// 	if j == nil {
// 		return []byte("null"), nil
// 	}
// 	return j, nil
// }

// UnmarshalJSON sets *j to a copy of data.
func (j *Bit16) UnmarshalJSON(data []byte) error {
	if j == nil {
		return fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = convBit16(string(data))
	return nil
}

func (j Bit16) Check(position int) bool {
	return j>>position&0x01 == 1
}
