package jsonld

import (
	"encoding/json"
	"errors"
	"reflect"
)

type Option int64

const NestedCall = iota

func extractFieldData(v reflect.Value, f func(v any, opts ...Option) ([]byte, error)) (any, error) {
	var fieldData interface{}

	switch v.Kind() {

	case reflect.Struct:
		var mapData map[string]interface{}
		bytes, err := f(v.Interface(), NestedCall)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(bytes, &mapData); err != nil {
			return nil, err
		}
		fieldData = mapData

	case reflect.Slice:
		var sliceData []interface{}
		for j := 0; j < v.Len(); j++ {
			var mapData map[string]interface{}
			bytes, err := f(v.Index(j).Interface(), NestedCall)
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &mapData); err != nil {
				return nil, err
			}
			sliceData = append(sliceData, mapData)
		}
		fieldData = sliceData

	default:
		fieldData = v.Interface()
	}

	return fieldData, nil
}

func hasNestedOpt(opts ...Option) bool {
	if len(opts) == 0 {
		return false
	}
	for _, opt := range opts {
		if opt == NestedCall {
			return true
		}
	}

	return false
}

func coverWithSquaredBrackets(in []byte) []byte {
	p := []byte("[")
	s := []byte("]")
	in = append(p, in...)
	in = append(in, s...)
	return in
}

func MarshalExpanded(v any, opts ...Option) ([]byte, error) {
	// 1. Convert struct to map[string]interface
	jsonldMap := map[string]interface{}{}
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct or  {
		return nil, errors.New("type unsupported by MarshalExpanded method")
	}

	if items.Kind() == reflect.Slice {
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if item.Kind() == reflect.Struct {
				v := reflect.Indirect(item)
				for j := 0; j < v.NumField(); j++ {
					fmt.Println(v.Type().Field(j).Name, v.Field(j).Interface())
				}
			}
		}
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		var fieldData interface{}
		var err error

		if !val.Field(i).CanInterface() {
			// Skip unexported fields
			continue
		}

		fieldData, err = extractFieldData(val.Field(i), MarshalExpanded)
		if err != nil {
			return nil, err
		}

		fieldTag := val.Type().Field(i).Tag.Get("jsonld")

		if fieldTag == "" || fieldTag == "-" {
			continue
		}

		// Fixed JSON LD Type definition
		if fieldTag == "@type" && val.Type().Field(i).Tag.Get("default") != "" {
			jsonldMap[fieldTag] = []string{
				val.Type().Field(i).Tag.Get("default"),
			}
			continue
		}

		jsonldMap[fieldTag] = []map[string]interface{}{
			{
				"@value": fieldData,
			},
		}

		if val.Field(i).Kind() == reflect.Struct {
			jsonldMap[fieldTag] = []interface{}{
				fieldData,
			}
		}

		if fieldTag == "@id" || val.Field(i).Kind() == reflect.Slice {
			jsonldMap[fieldTag] = fieldData
		}
	}

	// 2. Convert map[string]interface to []byte
	compactedDocJson, err := json.Marshal(jsonldMap)
	if err != nil {
		return nil, err
	}

	if !hasNestedOpt(opts...) {
		compactedDocJson = coverWithSquaredBrackets(compactedDocJson)
	}

	return compactedDocJson, nil
}

func MarshalCompacted(v any, opts ...Option) ([]byte, error) {
	// TODO refactor all Marshal methods
	// 1. Convert struct to map[string]interface
	jsonldMap := map[string]interface{}{}
	val := reflect.ValueOf(v)

	// if val.Kind() != reflect.Struct {
	// 	return nil, errors.New("type unsupported by MarshalCompacted method")
	// }

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		var fieldData interface{}
		var err error

		if !val.Field(i).CanInterface() {
			// Skip unexported fields
			continue
		}

		fieldData, err = extractFieldData(val.Field(i), MarshalCompacted)
		if err != nil {
			return nil, err
		}

		fieldTag := val.Type().Field(i).Tag.Get("jsonld")

		if fieldTag == "" || fieldTag == "-" {
			continue
		}

		// Fixed JSON LD Type definition
		if fieldTag == "@type" && val.Type().Field(i).Tag.Get("default") != "" {
			jsonldMap[fieldTag] = val.Type().Field(i).Tag.Get("default")
			continue
		}

		jsonldMap[fieldTag] = fieldData
	}

	// 2. Convert map[string]interface to []byte
	compactedDocJson, err := json.Marshal(jsonldMap)
	if err != nil {
		return nil, err
	}

	return compactedDocJson, nil
}

func MarshalContext(v any, c map[string]interface{}, opts ...Option) ([]byte, error) {
	return nil, errors.New("method not implemented yet")
}