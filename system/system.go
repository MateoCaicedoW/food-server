package system

import (
	"reflect"
	"strings"
)

// Function that generate a map[string][]string from a struct
func StructToMap(s interface{}) map[string][]string {
	m := make(map[string][]string)
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < v.NumField(); i++ {

		m[strings.ToLower(t.Field(i).Name)] = append(m[strings.ToLower(t.Field(i).Name)], v.Field(i).String())
	}
	return m
}
