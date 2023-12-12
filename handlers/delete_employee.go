package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewDeleteEmployee function will delete the employee
func NewDeleteEmployee(rt *runtime.Runtime) operations.DeleteEmployeeHandler {
	return &deleteEmployee{rt: rt}
}

type deleteEmployee struct {
	rt *runtime.Runtime
}

// Handle the delete employee request
func (d *deleteEmployee) Handle(params operations.DeleteEmployeeParams) middleware.Responder {
	if err := d.rt.Service().DeleteEmployee(params.ID); err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to delete employee: error[404]: %+v ", err)
			return operations.NewDeleteEmployeeNotFound()
		default:
			log().Errorf("ailed to delete employee: error[500]: %+v ", err)
			return operations.NewDeleteEmployeeInternalServerError()
		}
	}

	return operations.NewDeleteEmployeeNoContent()
}
