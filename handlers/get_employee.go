package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewGetEmployee handles a request for retrieving an employee
func NewGetEmployee(rt *runtime.Runtime) operations.GetEmployeeHandler {
	return &getEmployee{rt: rt}
}

type getEmployee struct {
	rt *runtime.Runtime
}

// Handle the get employee request
func (d *getEmployee) Handle(params operations.GetEmployeeParams) middleware.Responder {
	employee, err := d.rt.Service().GetEmployeeByID(params.ID)
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to get employee: error[404]: %+v ", err)
			return operations.NewGetEmployeeNotFound()
		default:
			log().Errorf("failed to get employee: error[500]: %+v ", err)
			return operations.NewGetEmployeeInternalServerError()
		}
	}

	return operations.NewGetEmployeeOK().WithPayload(toEmployeeGen(employee))
}
