package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/models"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewListEmployee handles a request for listing employee
func NewListEmployee(rt *runtime.Runtime) operations.ListEmployeesHandler {
	return &listEmployee{
		rt: rt,
	}
}

type listEmployee struct {
	rt *runtime.Runtime
}

// Handle the list employee request
func (l *listEmployee) Handle(params operations.ListEmployeesParams) middleware.Responder {
	employees, err := l.rt.Service().ListEmployee()
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NoContent):
			log().Errorf("failed to list employee: error[404]: %+v ", err)
			return operations.NewListEmployeesNoContent()
		default:
			log().Errorf("failed to list employee: error[500]: %+v ", err)
			return operations.NewListEmployeesInternalServerError()
		}
	}

	var payload []*models.Employee
	for _, employee := range employees {
		payload = append(payload, &models.Employee{
			ID:         employee.ID,
			Name:       employee.Name,
			Email:      employee.Email,
			Phone:      employee.Phone,
			Salary:     float32(employee.Salary),
			Benefits:   employee.Benefits,
			Title:      employee.Title,
			Team:       employee.Team,
			Outstation: employee.Outstation,
			StartDate:  strfmt.DateTime(employee.StartDate),
		})
	}

	return operations.NewListEmployeesOK().WithPayload(payload)
}
