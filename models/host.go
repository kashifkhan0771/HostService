package models

import "github.com/fatih/structs"

// Host Model
type Host struct {
	ID       string                 `structs:"id" json:"id" bson:"_id" db:"id"`
	IP       string                 `structs:"ip" json:"ip" bson:"ip" db:"ip"`
	Name     string                 `structs:"name" json:"name" bson:"name" db:"name"`
	Metadata map[string]interface{} `structs:"metadata" json:"metadata" bson:"metadata" db:"metadata"`
	Status   bool                   `structs:"status" json:"status" bson:"status" db:"status"`
}

// Map convert struct into map
func (h *Host) Map() map[string]interface{} {
	return structs.Map(h)
}

//Names return field names of Host model
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
