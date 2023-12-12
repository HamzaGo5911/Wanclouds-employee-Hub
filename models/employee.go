package models

import (
	"time"

	"github.com/fatih/structs"
)

type Employee struct {
	ID         string    `json:"id" structs:"id" bson:"_id" db:"id"`
	OfficeID   string    `json:"office_id" structs:"office_id" bson:"office_id"`
	HomeID     string    `json:"home_id" structs:"home_id" bson:"home_id"`
	Name       string    `json:"name" structs:"name" bson:"name" db:"name"`
	Email      string    `json:"email" structs:"email" bson:"email" db:"email"`
	Title      string    `json:"title" structs:"title" bson:"title" db:"title"`
	Phone      string    `json:"phone" structs:"phone" bson:"phone" db:"phone"`
	Team       string    `json:"team" structs:"team" bson:"team" db:"team"`
	StartDate  time.Time `json:"start_date" structs:"start_date" bson:"start_date" db:"start_date"`
	Salary     float64   `json:"salary" structs:"salary" bson:"salary" db:"salary"`
	Benefits   []string  `json:"benefits" structs:"benefits" bson:"benefits" db:"benefits"`
	Outstation bool      `json:"outstation" structs:"outstation" bson:"outstation" db:"outstation"`
}

// Map converts structs to a map representation
func (e *Employee) Map() map[string]interface{} {
	return structs.Map(e)
}

// Names returns the field names of Employee model
func (e *Employee) Names() []string {
	fields := structs.Fields(e)
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
