package models

import "github.com/fatih/structs"

type Office struct {
	ID          string `json:"id" structs:"id"`
	EmployeeID  string `json:"employee_id" structs:"employee_id"`
	OfficeRoom  string `json:"office_room" structs:"office_room"`
	MeetingRoom string `json:"meeting_room" structs:"meeting_room"`
	LunchArea   string `json:"lunch_area" structs:"lunch_area"`
	TeaArea     string `json:"tea_area" structs:"tea_area"`
	PlayingArea string `json:"playing_area" structs:"playing_area"`
}

// Map converts structs to a map representation
func (o *Office) Map() map[string]interface{} {
	return structs.Map(o)
}

// Names returns the field names of Office model
func (o *Office) Names() []string {
	fields := structs.Fields(o)
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
