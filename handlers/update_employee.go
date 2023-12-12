package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewUpdateEmployee handles request for updating employee
func NewUpdateEmployee(rt *runtime.Runtime) operations.UpdateEmployeeHandler {
	return &updateEmployee{
		rt: rt,
	}
}

type updateEmployee struct {
	rt *runtime.Runtime
}

// Handle the update employee request
func (d *updateEmployee) Handle(params operations.UpdateEmployeeParams) middleware.Responder {
	params.Employee.ID = params.ID

	_, err := d.rt.Service().UpdateEmployee(toEmployeeDomain(params.Employee))
	if err != nil {
		switch apiErr, ok := err.(*domainerr.APIError); ok {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to update employee: error[404]: %+v ", err)
			return operations.NewUpdateEmployeeNotFound()
		default:
			log().Errorf("failed to update empolyee: error[500]: %+v ", err)
			return operations.NewUpdateEmployeeInternalServerError()
		}
	}

	return operations.NewUpdateEmployeeOK().WithPayload(params.Employee)
}
