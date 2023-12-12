package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/models"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewListHome handles a request for listing home
func NewListHome(rt *runtime.Runtime) operations.ListHomesHandler {
	return &listHome{
		rt: rt,
	}
}

type listHome struct {
	rt *runtime.Runtime
}

// Handle the list home request
func (l *listHome) Handle(params operations.ListHomesParams) middleware.Responder {
	homes, err := l.rt.Service().ListHome()
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NoContent):
			log().Errorf("failed to list home: error[404]: %+v ", err)
			return operations.NewListHomesNoContent()
		default:
			log().Errorf("failed to list home: error[500]: %+v ", err)
			return operations.NewListHomesInternalServerError()
		}
	}

	var payload []*models.Home
	for _, home := range homes {
		payload = append(payload, &models.Home{
			ID:            home.ID,
			CommonRoom:    home.CommonRoom,
			DiningRoom:    home.DiningRoom,
			Dinner:        home.Dinner,
			DinnerRoom:    home.DinnerRoom,
			ResidencyRoom: home.ResidencyRoom,
			Tea:           home.Tea,
		})
	}

	return operations.NewListHomesOK().WithPayload(payload)
}
