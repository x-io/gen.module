package auth

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Index int
	Type  reflect.Kind
}

func filter(data interface{}, eye string) interface{} {
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	// log.Printf("data:[%v], eye:[%s]", data, eye)
	if rv.IsZero() {
		return data
	}

	switch rt.Kind() {
	case reflect.Array, reflect.Slice:
		if rv.Len() == 0 {
			return data
		}

		field := getField(rv.Index(0).Type(), eye)
		if len(field) > 0 {
			for i := 0; i < rv.Len(); i++ {
				change(rv.Index(i), field)
			}
		}
	case reflect.Struct:
		field := getField(rt, eye)
		if len(field) > 0 {
			change(rv, field)
		}
	case reflect.Ptr:
		field := getField(rt.Elem(), eye)
		if len(field) > 0 {
			change(rv.Elem(), field)
		}
	default:
		fmt.Println(rt.Kind())
	}

	return data
}

func getField(rt reflect.Type, eye string) []Field {
	var tag string
	ref := make([]Field, 0)

	//fmt.Println(rt.Kind())
	if rt.Kind() != reflect.Struct {
		return ref
	}

	for i := 0; i < rt.NumField(); i++ {
		tag = rt.Field(i).Tag.Get("eye")
		if tag != "" && (eye == "" || !strings.Contains(eye+",", tag+",")) {
			ref = append(ref, Field{i, rt.Field(i).Type.Kind()})
		}
	}

	return ref
}

func change(rv reflect.Value, field []Field) {
	for _, v := range field {
		//fmt.Println(v.Type)

		if rv.Field(v.Index).IsZero() {
			continue
		}

		switch v.Type {
		case reflect.Bool:
			rv.Field(v.Index).SetBool(false)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			rv.Field(v.Index).SetInt(0)
		case reflect.Float32, reflect.Float64:
			rv.Field(v.Index).SetFloat(0)
		case reflect.String:
			rv.Field(v.Index).SetString("")
		case reflect.Slice, reflect.Array:
			rv.Field(v.Index).SetBytes(nil)
		case reflect.Interface, reflect.Ptr:
			rv.Field(v.Index).SetPointer(nil)
		case reflect.Struct, reflect.Map:
			//rv.Field(v.Index).
		default:
		}
	}
}
