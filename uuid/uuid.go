package uuid

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type UUID struct {
	value uuid.UUID
}

func NewV7() UUID {
	return UUID{value: uuid.Must(uuid.NewV7())}
}

func New() UUID {
	return UUID{value: uuid.New()}
}

// MarshalJSON -> to JSON
func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value.String())
}

// UnmarshalJSON -> from JSON
func (u *UUID) UnmarshalJSON(data []byte) error {
	var jsonString string
	if err := json.Unmarshal(data, &jsonString); err != nil {
		return err
	}
	value, err := uuid.Parse(jsonString)
	if err != nil {
		return err
	}
	*u = UUID{value: value}
	return nil
}

// Scan --> From DB
func (u *UUID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	parsedUUID, err := uuid.FromBytes(bytes)
	*u = UUID{value: parsedUUID}
	return err
}

// Value -> TO DB
func (u UUID) Value() (driver.Value, error) {
	return u.value.MarshalBinary()
}

func (u UUID) String() string {
	return u.value.String()
}

func Parse(s string) (UUID, error) {
	u, err := uuid.Parse(s)
	return UUID{value: u}, err
}

func MustParse(s string) UUID {
	u := uuid.MustParse(s)
	return UUID{value: u}
}
