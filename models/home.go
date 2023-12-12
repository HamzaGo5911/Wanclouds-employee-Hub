package models

import "github.com/fatih/structs"

type Home struct {
	ID            string `json:"id" structs:"id" bson:"_id"`
	ResidencyRoom string `json:"residency_room" structs:"residency_room" bson:"residency_room"`
	Dinner        string `json:"dinner" structs:"dinner" bson:"dinner"`
	DinnerRoom    string `json:"dinner_room" structs:"dinner_room" bson:"dinner_room"`
	Tea           string `json:"tea" structs:"tea" bson:"tea"`
	DiningRoom    string `json:"dining_room" structs:"dining_room" bson:"dining_room"`
	CommonRoom    string `json:"common_room" structs:"common_room" bson:"common_room"`
}

// Map converts structs to a map representation
func (h *Home) Map() map[string]interface{} {
	return structs.Map(h)
}

// Names returns the field names of Home model
func (h *Home) Names() []string {
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
