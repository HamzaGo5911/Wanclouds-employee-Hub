package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// NewAddHome handles request for saving home
func NewAddHome(rt *runtime.Runtime) operations.AddHomeHandler {
	return &addHome{rt: rt}
}

type addHome struct {
	rt *runtime.Runtime
}

// Handle the add home request
func (d *addHome) Handle(params operations.AddHomeParams) middleware.Responder {
	newHome := models.Home{
		ID:            params.Home.ID,
		CommonRoom:    params.Home.CommonRoom,
		DiningRoom:    params.Home.DiningRoom,
		Dinner:        params.Home.Dinner,
		DinnerRoom:    params.Home.DinnerRoom,
		Tea:           params.Home.Tea,
		ResidencyRoom: params.Home.ResidencyRoom,
	}

	homeID, err := d.rt.Service().AddHome(&newHome)
	if err != nil {
		var apiErr *domainerr.APIError
		if errors.As(err, &apiErr) {
			switch {
			case apiErr.IsError(domainerr.Conflict):
				log().Errorf("Failed to add home: Error [409]: %s", err.Error())
				return operations.NewAddHomeConflict()
			case apiErr.IsError(domainerr.BadRequest):
				log().Errorf("Failed to add home: Error [400]: %s", err.Error())
				return operations.NewAddHomeBadRequest()
			default:
				log().Errorf("Failed to add home: Error [500]: %s", err.Error())
				return operations.NewAddHomeInternalServerError()
			}
		}
	}

	log().Infof("Home added successfully")

	params.Home.ID = homeID
	return operations.NewAddHomeCreated().WithPayload(params.Home)

}
