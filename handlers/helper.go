package handlers

import (
	"github.com/go-openapi/strfmt"

	genmodels "github.com/HamzaGo5911/wanclouds-employee-hub/gen/models"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

func toEmployeeGen(employee *models.Employee) *genmodels.Employee {
	return &genmodels.Employee{
		ID:         employee.ID,
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Team:       employee.Team,
		Title:      employee.Title,
		StartDate:  strfmt.DateTime(employee.StartDate),
		Salary:     float32(employee.Salary),
		Benefits:   employee.Benefits,
		Outstation: employee.Outstation,
	}
}

func toOfficeGen(office *models.Office) *genmodels.Office {
	return &genmodels.Office{
		ID:          office.ID,
		Room:        office.Room,
		MeetingRoom: office.MeetingRoom,
		LunchArea:   office.LunchArea,
		TeaArea:     office.TeaArea,
		PlayingArea: office.PlayingArea,
	}
}

func toHomeGen(home *models.Home) *genmodels.Home {
	return &genmodels.Home{
		ID:            home.ID,
		CommonRoom:    home.CommonRoom,
		DinnerRoom:    home.DinnerRoom,
		DiningRoom:    home.DiningRoom,
		Dinner:        home.Dinner,
		ResidencyRoom: home.ResidencyRoom,
	}
}

func toEmployeeDomain(employee *genmodels.Employee) *models.Employee {
	return &models.Employee{
		ID:       employee.ID,
		HomeID:   employee.HomeID,
		OfficeID: employee.OfficeID,
		Name:     employee.Name,
		Email:    employee.Email,
		Phone:    employee.Phone,
		Team:     employee.Team,
		Title:    employee.Title,
		Salary:   float64(employee.Salary),
		Benefits: employee.Benefits,
	}
}

func toHomeDomain(home *genmodels.Home) *models.Home {
	return &models.Home{
		ID:            home.ID,
		ResidencyRoom: home.ResidencyRoom,
		Dinner:        home.Dinner,
		DiningRoom:    home.DiningRoom,
		DinnerRoom:    home.DinnerRoom,
		CommonRoom:    home.CommonRoom,
	}
}

func toOfficeDomain(office *genmodels.Office) *models.Office {
	return &models.Office{
		ID:          office.ID,
		Name:        office.Name,
		LunchArea:   office.LunchArea,
		Room:        office.Room,
		MeetingRoom: office.MeetingRoom,
		PlayingArea: office.PlayingArea,
		TeaArea:     office.TeaArea,
	}
}
