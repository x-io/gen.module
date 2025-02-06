package dbs

import (
	"database/sql"
	"log"
	"strings"
)

type DBError struct {
	ECode   int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Error Error
func (e *DBError) Error() string {
	return e.Message
}

// Code Code
func (e *DBError) Code() int {
	return e.ECode
}

// Error Error
func Error(err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return err
	}
	e := err.Error()
	if strings.HasPrefix(e, "write tcp") {
		return &DBError{Message: "服务异常"}
	}
	return &DBError{Message: e}
}

// Error Error
func Errorf(message string, a ...interface{}) error {

	if len(a) > 0 {
		if v, ok := a[0].(error); ok {
			if v == sql.ErrNoRows {
				log.Println(v)
			} else {
				log.Println(v)
			}
		} else {
			log.Println(a...)
		}
	}

	return &DBError{Message: message}
}

// ErrorCode NewCode
func ErrorCode(code int, data interface{}) error {
	if v, ok := data.(string); ok {
		return &DBError{ECode: code, Message: v}
	}

	return &DBError{ECode: code, Data: data}
}
