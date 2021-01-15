package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
)

type metadata map[string]interface{}

// Host Model.
type Host struct {
	ID       string   `structs:"id" json:"id" bson:"_id" db:"id"`
	IP       string   `structs:"ip" json:"ip" bson:"ip" db:"ip"`
	Name     string   `structs:"name" json:"name" bson:"name" db:"name"`
	Metadata metadata `structs:"metadata" json:"metadata" bson:"metadata" db:"metadata"`
	Status   bool     `structs:"status" json:"status" bson:"status" db:"status"`
}

func (md metadata) Value() (driver.Value, error) {
	j, err := json.Marshal(md)

	return j, err
}

func (md *metadata) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		fmt.Printf("Type assertion .([]byte) failed.")
	}

	var i interface{}
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*md, ok = i.(map[string]interface{})
	if !ok {
		fmt.Printf("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

// Map convert struct into map.
func (h *Host) Map() map[string]interface{} {
	return structs.Map(h)
}

// Names return field names of Host model.
func (h *Host) Names() []string {
	fields := structs.Fields(h)

	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
