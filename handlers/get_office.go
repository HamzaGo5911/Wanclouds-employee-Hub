package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewGetOffice handles a request for retrieving an office
func NewGetOffice(rt *runtime.Runtime) operations.GetOfficeHandler {
	return &getOffice{rt: rt}
}

type getOffice struct {
	rt *runtime.Runtime
}

// Handle the get employee request
func (d *getOffice) Handle(params operations.GetOfficeParams) middleware.Responder {
	office, err := d.rt.Service().GetOfficeByID(params.ID)
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to get office: error[404]: %+v ", err)
			return operations.NewGetOfficeNotFound()
		default:
			log().Errorf("failed to get office: error[500]: %+v ", err)
			return operations.NewGetOfficeInternalServerError()
		}
	}

	return operations.NewGetOfficeOK().WithPayload(toOfficeGen(office))
}
