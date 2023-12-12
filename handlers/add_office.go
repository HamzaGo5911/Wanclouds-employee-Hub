package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// NewAddOffice handles request for saving office
func NewAddOffice(rt *runtime.Runtime) operations.AddOfficeHandler {
	return &addOffice{rt: rt}
}

type addOffice struct {
	rt *runtime.Runtime
}

// Handle the add office request
func (d *addOffice) Handle(params operations.AddOfficeParams) middleware.Responder {
	newOffice := models.Office{
		ID:          params.Office.ID,
		Name:        params.Office.Name,
		Room:        params.Office.Room,
		MeetingRoom: params.Office.MeetingRoom,
		LunchArea:   params.Office.LunchArea,
		PlayingArea: params.Office.PlayingArea,
		TeaArea:     params.Office.TeaArea,
	}

	officeID, err := d.rt.Service().AddOffice(&newOffice)
	if err != nil {
		var apiErr *domainerr.APIError
		if errors.As(err, &apiErr) {
			switch {
			case apiErr.IsError(domainerr.Conflict):
				log().Errorf("Failed to add office: Error [409]: %s", err.Error())
				return operations.NewAddOfficeConflict()
			case apiErr.IsError(domainerr.BadRequest):
				log().Errorf("Failed to add office: Error [400]: %s", err.Error())
				return operations.NewAddOfficeBadRequest()
			default:
				log().Errorf("Failed to add office: Error [500]: %s", err.Error())
				return operations.NewAddOfficeInternalServerError()
			}
		}
	}

	log().Infof("Office added successfully")

	params.Office.ID = officeID
	return operations.NewAddOfficeCreated().WithPayload(params.Office)

}
