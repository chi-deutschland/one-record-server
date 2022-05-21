package jsonld

import (
	"errors"
	"github.com/go-ap/jsonld"
)

func UnmarshalExpanded(data []byte, v any) error {
	return errors.New("method not implemented yet")
}

func UnmarshalCompacted(data []byte, v any) error {
	err := jsonld.Unmarshal(data, v)
	if err != nil {
		return err
	}
	// TODO add check of @type value with default @type-value in struct
	return nil
}

func UnmarshalContext(data []byte, c map[string]interface{}, v any) error {
	return errors.New("method not implemented yet")
}