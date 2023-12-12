package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewGetHome handles a request for retrieving a home
func NewGetHome(rt *runtime.Runtime) operations.GetHomeHandler {
	return &getHome{rt: rt}
}

type getHome struct {
	rt *runtime.Runtime
}

// Handle the get home request
func (d *getHome) Handle(params operations.GetHomeParams) middleware.Responder {
	home, err := d.rt.Service().GetHomeByID(params.ID)
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NotFound):
			log().Errorf("failed to get home: error[404]: %+v ", err)
			return operations.NewGetHomeNotFound()
		default:
			log().Errorf("failed to get home: error[500]: %+v ", err)
			return operations.NewGetHomeInternalServerError()
		}
	}

	return operations.NewGetHomeOK().WithPayload(toHomeGen(home))
}
