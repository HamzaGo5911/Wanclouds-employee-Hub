package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewDeleteHome function will delete the home
func NewDeleteHome(rt *runtime.Runtime) operations.DeleteHomeHandler {
	return &deleteHome{rt: rt}
}

type deleteHome struct {
	rt *runtime.Runtime
}

// Handle the delete home request
func (d *deleteHome) Handle(params operations.DeleteHomeParams) middleware.Responder {
	if err := d.rt.Service().DeleteHome(params.ID); err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to delete home: error[404]: %+v ", err)
			return operations.NewDeleteHomeNotFound()
		default:
			log().Errorf("ailed to delete home: error[500]: %+v ", err)
			return operations.NewDeleteHomeInternalServerError()
		}
	}

	return operations.NewDeleteHomeNoContent()
}
