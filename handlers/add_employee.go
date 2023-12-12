package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// NewAddEmployee handles request for saving employee
func NewAddEmployee(rt *runtime.Runtime) operations.AddEmployeeHandler {
	return &addEmployee{rt: rt}
}

type addEmployee struct {
	rt *runtime.Runtime
}

// Handle the add employee request
func (d *addEmployee) Handle(params operations.AddEmployeeParams) middleware.Responder {
	newEmployee := models.Employee{
		ID:         params.Employee.ID,
		OfficeID:   params.Employee.OfficeID,
		HomeID:     params.Employee.HomeID,
		Name:       params.Employee.Name,
		Email:      params.Employee.Email,
		Title:      params.Employee.Title,
		Phone:      params.Employee.Phone,
		Team:       params.Employee.Team,
		Outstation: params.Employee.Outstation,
		StartDate:  time.Time(params.Employee.StartDate),
		Salary:     float64(params.Employee.Salary),
		Benefits:   params.Employee.Benefits,
	}

	employeeID, err := d.rt.Service().AddEmployee(&newEmployee)
	if err != nil {
		var apiErr *domainerr.APIError
		if errors.As(err, &apiErr) {
			switch {
			case apiErr.IsError(domainerr.Conflict):
				log().Errorf("Failed to add employee: Error [409]: %s", err.Error())
				return operations.NewAddEmployeeConflict()
			case apiErr.IsError(domainerr.NotFound):
				log().Errorf("Failed to add employee: Error [400]: %s", err.Error())
				return operations.NewAddEmployeeBadRequest()
			default:
				log().Errorf("Failed to add employee: Error [500]: %s", err.Error())
				return operations.NewAddEmployeeInternalServerError()
			}
		}
	}

	log().Infof("Employee added successfully")

	params.Employee.ID = employeeID
	return operations.NewAddEmployeeCreated().WithPayload(params.Employee)

}
