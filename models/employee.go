package models

import (
	"time"

	"github.com/fatih/structs"
)

type Employee struct {
	ID        string    `json:"id" structs:"id"`
	Name      string    `json:"name" structs:"name"`
	Email     string    `json:"email" structs:"email"`
	Phone     string    `json:"phone" structs:"phone"`
	Team      string    `json:"team" structs:"team"`
	JobTitle  string    `json:"job_title" structs:"job_title"`
	StartDate time.Time `json:"start_date" structs:"start_date"`
	Salary    float64   `json:"salary" structs:"salary"`
	Benefits  []string  `json:"benefits" structs:"benefits"`
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
