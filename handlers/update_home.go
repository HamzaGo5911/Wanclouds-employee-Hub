package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewUpdateHome handles request for updating home
func NewUpdateHome(rt *runtime.Runtime) operations.UpdateHomeHandler {
	return &updateHome{
		rt: rt,
	}
}

type updateHome struct {
	rt *runtime.Runtime
}

// Handle the update home request
func (d *updateHome) Handle(params operations.UpdateHomeParams) middleware.Responder {
	params.Home.ID = params.ID

	_, err := d.rt.Service().UpdateHome(toHomeDomain(params.Home))
	if err != nil {
		switch apiErr, ok := err.(*domainerr.APIError); ok {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to update home: error[404]: %+v ", err)
			return operations.NewUpdateHomeNotFound()
		default:
			log().Errorf("failed to home: error[500]: %+v ", err)
			return operations.NewUpdateHomeInternalServerError()
		}
	}

	return operations.NewUpdateHomeOK().WithPayload(params.Home)
}
