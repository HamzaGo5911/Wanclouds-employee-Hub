package models

import (
	"reflect"
	"testing"
)

func TestOffice_Map(t *testing.T) {
	office := &Office{
		ID:          "1",
		Name:        "wanclouds",
		Room:        "Room 405",
		MeetingRoom: "Room 206",
		LunchArea:   "gallery",
		TeaArea:     "Kitchen",
		PlayingArea: "Recreation Room",
	}

	officeMap := office.Map()

	expectedMap := map[string]interface{}{
		"id":           "1",
		"name":         "wanclouds",
		"room":         "Room 405",
		"meeting_room": "Room 206",
		"lunch_area":   "gallery",
		"tea_area":     "Kitchen",
		"playing_area": "Recreation Room",
	}

	if !reflect.DeepEqual(officeMap, expectedMap) {
		t.Errorf("Expected map %v, but got %v", expectedMap, officeMap)
	}
}

func TestOffice_Names(t *testing.T) {
	office := &Office{}

	fieldNames := office.Names()

	expectedNames := []string{
		"id",
		"name",
		"room",
		"meeting_room",
		"lunch_area",
		"tea_area",
		"playing_area",
	}

	if !reflect.DeepEqual(fieldNames, expectedNames) {
		t.Errorf("Expected field names %v, but got %v", expectedNames, fieldNames)
	}
}
