package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/models"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// NewListOffice handles a request for listing office
func NewListOffice(rt *runtime.Runtime) operations.ListOfficesHandler {
	return &listOffice{
		rt: rt,
	}
}

type listOffice struct {
	rt *runtime.Runtime
}

// Handle the list office request
func (l *listOffice) Handle(params operations.ListOfficesParams) middleware.Responder {
	offices, err := l.rt.Service().ListOffice()
	if err != nil {
		switch apiErr := err.(*domainerr.APIError); {
		case apiErr.IsError(domainerr.NoContent):
			log().Errorf("failed to list office: error[404]: %+v ", err)
			return operations.NewListOfficesNoContent()
		default:
			log().Errorf("failed to list office: error[500]: %+v ", err)
			return operations.NewListOfficesInternalServerError()
		}
	}

	var payload []*models.Office
	for _, office := range offices {
		payload = append(payload, &models.Office{
			ID:          office.ID,
			LunchArea:   office.LunchArea,
			MeetingRoom: office.MeetingRoom,
			Room:        office.Room,
			TeaArea:     office.TeaArea,
		})
	}

	return operations.NewListOfficesOK().WithPayload(payload)
}
