package service

import (
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// AddHome adds home into the database
func (s *Service) AddHome(home *models.Home) (string, error) {
	existingHomes, err := s.db.ListHome()
	if err != nil {
		return "", err
	}

	if len(existingHomes) >= 2 {
		return "", domainerr.NewAPIError(domainerr.Conflict, "home already exists")
	}

	addedHomeID, err := s.db.AddHome(home)
	if err != nil {
		return "", err
	}

	return addedHomeID, nil
}

// GetHomeByID retrieves a home by its ID
func (s *Service) GetHomeByID(id string) (*models.Home, error) {
	home, err := s.db.GetHomeByID(id)
	if err != nil {
		return nil, err
	}

	return home, nil
}

// DeleteHome deletes a home by its ID
func (s *Service) DeleteHome(id string) error {
	if _, err := s.db.GetHomeByID(id); err != nil {
		return err
	}

	return s.db.DeleteHome(id)
}

// ListHome retrieves a list of all home
func (s *Service) ListHome() ([]*models.Home, error) {
	home, err := s.db.ListHome()
	if err != nil {
		return nil, err
	}

	return home, nil
}

// UpdateHome updates an existing home with the provided information.
func (s *Service) UpdateHome(home *models.Home) (string, error) {
	homeID, err := s.db.GetHomeByID(home.ID)
	if err != nil {
		return "", err
	}

	if home.ID != homeID.ID {
		return "", domainerr.NewAPIError(domainerr.NotFound, "home with ID does not exist")

	}

	existingHome, err := s.db.AddHome(home)
	if err != nil {
		return "", err
	}

	return existingHome, nil
}
