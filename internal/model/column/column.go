package column

import (
	"database/sql/driver"
	"encoding/json"
)

type SArr []string

func (c SArr) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *SArr) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &c)
}
