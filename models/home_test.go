package models

import (
	"reflect"
	"testing"
)

func TestHome_Map(t *testing.T) {
	home := &Home{
		EmployeeID:    "123",
		ResidencyRoom: "Room A",
		Dinner:        "Home-cooked",
		DinnerRoom:    "Kitchen",
		DiningRoom:    "Dining Hall",
		CommonRoom:    "Living Room",
	}

	homeMap := home.Map()

	expectedMap := map[string]interface{}{
		"employee_id":    "123",
		"residency_room": "Room A",
		"dinner":         "Home-cooked",
		"dinner_room":    "Kitchen",
		"dining_room":    "Dining Hall",
		"common_room":    "Living Room",
	}

	if !reflect.DeepEqual(homeMap, expectedMap) {
		t.Errorf("Expected map %v, but got %v", expectedMap, homeMap)
	}
}

func TestHome_Names(t *testing.T) {
	home := &Home{}

	fieldNames := home.Names()

	expectedNames := []string{
		"employee_id",
		"residency_room",
		"dinner",
		"dinner_room",
		"dining_room",
		"common_room",
	}

	if !reflect.DeepEqual(fieldNames, expectedNames) {
		t.Errorf("Expected field names %v, but got %v", expectedNames, fieldNames)
	}
}
