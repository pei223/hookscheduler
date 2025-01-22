package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// SQLBoilerでjsonbに対応するカスタム型
type JSONB map[string]any

// Scan implements the sql.Scanner interface for JSONB
func (j *JSONB) Scan(src interface{}) error {
	if src == nil {
		*j = nil
		return nil
	}

	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		return errors.New("unsupported type for JSONB")
	}
}

// Value implements the driver.Valuer interface for JSONB
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
