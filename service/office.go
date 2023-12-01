package service

import (
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// AddOffice adds office into the database
func (s *Service) AddOffice(office *models.Office) (string, error) {
	existingOffice, err := s.db.GetOfficeByID(office.ID)
	if err != nil {
		return "", err
	}

	offices, err := s.db.AddOffice(existingOffice)
	if err != nil {
		return "", err
	}

	return offices, nil
}

// GetOfficeByID retrieves a office by its ID
func (s *Service) GetOfficeByID(id string) (*models.Office, error) {
	office, err := s.db.GetOfficeByID(id)
	if err != nil {
		return nil, err
	}

	return office, nil
}

// DeleteOffice deletes a office by its ID
func (s *Service) DeleteOffice(id string) error {
	if _, err := s.db.GetOfficeByID(id); err != nil {
		return err
	}

	return s.db.DeleteOffice(id)
}

// ListOffice retrieves a list of all office
func (s *Service) ListOffice() ([]*models.Office, error) {
	office, err := s.db.ListOffice()
	if err != nil {
		return nil, err
	}

	return office, nil
}

// UpdateOffice updates an existing office with the provided information.
func (s *Service) UpdateOffice(office *models.Office) (string, error) {
	officeID, err := s.db.GetEmployeeByID(office.ID)
	if err != nil {
		return "", err
	}

	if office.ID != officeID.ID {
		return "", domainerr.NewAPIError(domainerr.NotFound, "office with ID does not exist")

	}

	existingOffice, err := s.db.AddOffice(office)
	if err != nil {
		return "", err
	}

	return existingOffice, nil
}
