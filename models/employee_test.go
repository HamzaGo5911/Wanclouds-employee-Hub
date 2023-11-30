package models

import (
	"reflect"
	"testing"
	"time"
)

func TestEmployee_Map(t *testing.T) {
	employee := &Employee{
		ID:        "123",
		Name:      "engineer hamza",
		Email:     "engrhamza@gmail.com",
		Phone:     "123-456-7890",
		Team:      "Backend Golang",
		JobTitle:  "Software Engineer",
		StartDate: time.Now(),
		Salary:    20000.0,
		Benefits:  []string{"Health Insurance", "Paid Time Off"},
	}

	employeeMap := employee.Map()

	expectedMap := map[string]interface{}{
		"id":         "123",
		"name":       "engineer hamza",
		"email":      "engrhamza@gmail.com",
		"phone":      "123-456-7890",
		"team":       "Backend Golang",
		"job_title":  "Software Engineer",
		"start_date": employee.StartDate,
		"salary":     20000.0,
		"benefits":   []string{"Health Insurance", "Paid Time Off"},
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
		"name",
		"email",
		"phone",
		"team",
		"job_title",
		"start_date",
		"salary",
		"benefits",
	}

	if !reflect.DeepEqual(fieldNames, expectedNames) {
		t.Errorf("Expected field names %v, but got %v", expectedNames, fieldNames)
	}
}
