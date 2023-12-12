package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewUpdateOffice handles request for updating office
func NewUpdateOffice(rt *runtime.Runtime) operations.UpdateOfficeHandler {
	return &updateOffice{
		rt: rt,
	}
}

type updateOffice struct {
	rt *runtime.Runtime
}

// Handle the update office request
func (d *updateOffice) Handle(params operations.UpdateOfficeParams) middleware.Responder {
	params.Office.ID = params.ID

	_, err := d.rt.Service().UpdateOffice(toOfficeDomain(params.Office))
	if err != nil {
		switch apiErr, ok := err.(*domainerr.APIError); ok {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to update office: error[404]: %+v ", err)
			return operations.NewUpdateOfficeNotFound()
		default:
			log().Errorf("failed to update office: error[500]: %+v ", err)
			return operations.NewUpdateOfficeInternalServerError()
		}
	}

	return operations.NewUpdateOfficeOK().WithPayload(params.Office)
}
