package column

import (
	"database/sql/driver"
	"encoding/json"
)

// SA string array type
type SA []string

func (c SA) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *SA) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &c)
}
