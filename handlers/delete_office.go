package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewDeleteOffice function will delete the office
func NewDeleteOffice(rt *runtime.Runtime) operations.DeleteOfficeHandler {
	return &deleteOffice{rt: rt}
}

type deleteOffice struct {
	rt *runtime.Runtime
}

// Handle the delete office request
func (d *deleteOffice) Handle(params operations.DeleteOfficeParams) middleware.Responder {
	if err := d.rt.Service().DeleteOffice(params.ID); err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to delete office: error[404]: %+v ", err)
			return operations.NewDeleteOfficeNotFound()
		default:
			log().Errorf("ailed to delete office: error[500]: %+v ", err)
			return operations.NewDeleteOfficeInternalServerError()
		}
	}

	return operations.NewDeleteOfficeNoContent()
}
