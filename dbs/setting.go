package dbs

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Setting map[string]interface{}

// Value Value
func (j Setting) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan Scan
func (j *Setting) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	var data []byte

	switch v := src.(type) {
	case string:
		data = []byte(fmt.Sprintf(`"%s"`, v))
	case []byte:
		if bytes.Equal(v, []byte{123, 125}) || bytes.Equal(v, []byte{91, 93}) {
			return nil
		}
		data = v
	}

	return json.Unmarshal(data, j)
}

func (s Setting) GetBool(name string) bool {
	if v, ok := s[name]; ok {
		if vv, ok := v.(bool); ok {
			return vv
		}
	}

	return false
}

func (s Setting) GetString(name string) string {
	if v, ok := s[name]; ok {
		if vv, ok := v.(string); ok {
			return vv
		}
	}

	return ""
}

func (s Setting) Get(name string) interface{} {
	// if v, ok := s[name]; ok {
	// 	if vv, ok := v.(string); ok {
	// 		return vv
	// 	}
	// }

	return s[name]
}

func (s Setting) GetMap(name ...string) interface{} {
	data := s

	for _, key := range name {
		if v, ok := data[key].(map[string]interface{}); ok {
			data = v
		} else {
			return data[key]
		}
	}

	return nil
}
