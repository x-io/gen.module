package dbs

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// Bit int
type Bit int

// Value Value
func (j Bit) Value() (driver.Value, error) {
	//log.Println(int64(j), fmt.Sprintf("%08b", j))
	return fmt.Sprintf("%08b", j), nil
}

// Scan Scan
func (j *Bit) Scan(src interface{}) error {
	if src == nil {
		*j = 0
		return nil
	}
	switch v := src.(type) {
	case []byte:
		*j = conv(string(v))
	case int64:
		*j = Bit(v)
	case string:
		*j = conv(v)
	}

	return nil
}

func conv(str string) Bit {
	str = strings.ReplaceAll(str, "\"", "")

	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		i, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	return Bit(i)
}

// // MarshalJSON returns j as the JSON encoding of j.
// func (J Bool) MarshalJSON() ([]byte, error) {
// 	if j == nil {
// 		return []byte("null"), nil
// 	}
// 	return j, nil
// }

// UnmarshalJSON sets *j to a copy of data.
func (j *Bit) UnmarshalJSON(data []byte) error {
	if j == nil {
		return fmt.Errorf("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = conv(string(data))
	return nil
}

func (j Bit) Check(position int) bool {
	return j>>position&0x01 == 1
}
