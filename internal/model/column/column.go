package column

import (
	"database/sql/driver"
	"encoding/json"
)

// SA string array type
type SA []string

func (c SA) Value() (driver.Value, error) {
	if len(c) > 0 && c[0] == "" {
		return nil, nil
	} else {
		return json.Marshal(c)
	}
}

func (c *SA) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &c)
}
