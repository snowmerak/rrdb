package rrdb

import (
	"reflect"
	"sync"
)

var dataache = sync.Map{}

// GetColumnsOf returns the columns of data.
// If the data is not a struct, it will return nil.
// Column type must be *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64, *string, *bool, *time.Time, *[]byte.
func GetColumnsOf(data any) []string {
	if v, ok := dataache.Load(reflect.TypeOf(data)); ok {
		return v.([]string)
	}

	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return nil
	}

	columns := make([]string, dataType.NumField())
	for i := 0; i < dataType.NumField(); i++ {
		if dataType.Field(i).Type.Kind() == reflect.Pointer {
			columns[i] = dataType.Field(i).Name
		}
	}

	dataache.Store(data, columns)
	return columns
}

// GetColumnsOfDataByCache returns the columns of data.
// If the data is not a struct, it will return nil.
func GetNotNilColumnsOf(data any) []string {
	if v, ok := dataache.Load(reflect.TypeOf(data)); ok {
		return v.([]string)
	}

	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return nil
	}

	columns := make([]string, 0, dataType.NumField())
	for i := 0; i < dataType.NumField(); i++ {
		if dataType.Field(i).Type.Kind() == reflect.Ptr && !reflect.ValueOf(data).Field(i).IsNil() {
			columns = append(columns, dataType.Field(i).Name)
		}
	}

	dataache.Store(data, columns)
	return columns
}
