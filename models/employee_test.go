package models

import (
	"reflect"
	"testing"
	"time"
)

func TestEmployee_Map(t *testing.T) {
	employee := &Employee{
		ID:         "123",
		OfficeID:   "1",
		HomeID:     "2",
		Name:       "engineer hamza",
		Email:      "engrhamza@gmail.com",
		Phone:      "123-456-7890",
		Team:       "Backend Golang",
		Title:      "Software Engineer",
		StartDate:  time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		Salary:     20000.0,
		Benefits:   []string{"Health Insurance", "Paid Time Off"},
		Outstation: true,
	}

	employeeMap := employee.Map()

	expectedMap := map[string]interface{}{
		"id":         "123",
		"office_id":  "1",
		"home_id":    "2",
		"name":       "engineer hamza",
		"email":      "engrhamza@gmail.com",
		"phone":      "123-456-7890",
		"team":       "Backend Golang",
		"title":      "Software Engineer",
		"start_date": time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		"salary":     20000.0,
		"benefits":   []string{"Health Insurance", "Paid Time Off"},
		"outstation": true,
	}

	if !reflect.DeepEqual(employeeMap, expectedMap) {
		t.Errorf("Expected map %v, but got %v", expectedMap, employeeMap)
	}
}

func TestEmployee_Names(t *testing.T) {
	employee := &Employee{}

	fieldNames := employee.Names()

	expectedNames := []string{
		"id",
		"office_id",
		"home_id",
		"name",
		"email",
		"title",
		"phone",
		"team",
		"start_date",
		"salary",
		"benefits",
		"outstation",
	}

	if !reflect.DeepEqual(fieldNames, expectedNames) {
		t.Errorf("Expected field names %v, but got %v", expectedNames, fieldNames)

	}
}
