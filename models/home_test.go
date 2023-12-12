package models

import (
	"reflect"
	"testing"
)

func TestHome_Map(t *testing.T) {
	home := &Home{
		ID:            "2",
		ResidencyRoom: "Room A",
		Dinner:        "Home-cooked",
		DinnerRoom:    "Kitchen",
		DiningRoom:    "Dining Hall",
		Tea:           "karak chai",
		CommonRoom:    "Living Room",
	}

	homeMap := home.Map()

	expectedMap := map[string]interface{}{
		"id":             "2",
		"residency_room": "Room A",
		"dinner":         "Home-cooked",
		"dinner_room":    "Kitchen",
		"dining_room":    "Dining Hall",
		"tea":            "karak chai",
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
		"id",
		"residency_room",
		"dinner",
		"dinner_room",
		"tea",
		"dining_room",
		"common_room",
	}

	if !reflect.DeepEqual(fieldNames, expectedNames) {
		t.Errorf("Expected field names %v, but got %v", expectedNames, fieldNames)
	}
}
