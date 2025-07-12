package json

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Container[T any] struct {
	RawValue T
}

func NewContainer[T any](value T) Container[T] {
	return Container[T]{RawValue: value}
}

// Scan --> From DB
func (j *Container[T]) Scan(value interface{}) error {
	val, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal string value:", value))
	}

	return json.Unmarshal(val, &j.RawValue)
}

// Value -> TO DB
func (j Container[T]) Value() (driver.Value, error) {
	val, err := json.Marshal(&j.RawValue)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// encoding.BinaryMarshaler and BinaryUnmarshaler
func (j Container[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(j.RawValue)
}

func (j *Container[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &j.RawValue)
}
