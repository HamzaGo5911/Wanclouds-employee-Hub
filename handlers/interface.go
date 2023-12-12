package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/HamzaGo5911/wanclouds-employee-hub"
	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.WancloudsEmpolyeeHubAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewWancloudsEmpolyeeHubAPI(spec)

	// employee handlers
	handler.AddEmployeeHandler = NewAddEmployee(rt)
	handler.GetEmployeeHandler = NewGetEmployee(rt)
	handler.DeleteEmployeeHandler = NewDeleteEmployee(rt)
	handler.ListEmployeesHandler = NewListEmployee(rt)
	handler.UpdateEmployeeHandler = NewUpdateEmployee(rt)

	// office handlers
	handler.AddOfficeHandler = NewAddOffice(rt)
	handler.DeleteOfficeHandler = NewDeleteOffice(rt)
	handler.GetOfficeHandler = NewGetOffice(rt)
	handler.ListOfficesHandler = NewListOffice(rt)
	handler.UpdateOfficeHandler = NewUpdateOffice(rt)

	// home handlers
	handler.AddHomeHandler = NewAddHome(rt)
	handler.DeleteHomeHandler = NewDeleteHome(rt)
	handler.GetHomeHandler = NewGetHome(rt)
	handler.ListHomesHandler = NewListHome(rt)
	handler.UpdateHomeHandler = NewUpdateHome(rt)

	return handler
}
