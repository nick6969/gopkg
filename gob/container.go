package gob

import (
	"bytes"
	"encoding/gob"
)

type Container[T any] struct {
	RawValue T
}

// encoding.BinaryMarshaler and BinaryUnmarshaler
func (c Container[T]) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(c.RawValue); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Container[T]) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&c.RawValue); err != nil {
		return err
	}
	return nil
}
