package service

import (
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// AddEmployee adds an employee into the database based on outstation status
func (s *Service) AddEmployee(employee *models.Employee) (string, error) {
	Office, err := s.db.GetOfficeByID(employee.OfficeID)
	if err != nil {
		return "", err
	}

	if Office.ID != employee.OfficeID {
		return "", domainerr.NewAPIError(domainerr.NotFound, "office does not exist")
	}

	existingEmployee, err := s.db.ListEmployee()
	if err != nil {
		return "", err
	}

	for _, emp := range existingEmployee {
		if emp.Email == employee.Email {
			return "", domainerr.NewAPIError(domainerr.Conflict, "an employee with the same email already exists")
		}
	}

	if !employee.Outstation {
		employee.HomeID = ""
	}

	if employee.Outstation {
		home, err := s.db.GetHomeByID(employee.HomeID)
		if err != nil {
			return "", err
		}

		if employee.HomeID == "" || employee.HomeID != home.ID {
			return "", domainerr.NewAPIError(domainerr.NotFound, "invalid HomeID for outstation employee")
		}
	}

	employeeID, err := s.db.AddEmployee(employee)
	if err != nil {
		return "", err
	}

	return employeeID, nil
}

// GetEmployeeByID retrieves a employee by its ID
func (s *Service) GetEmployeeByID(id string) (*models.Employee, error) {
	employee, err := s.db.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// DeleteEmployee deletes a employee by its ID
func (s *Service) DeleteEmployee(id string) error {
	if _, err := s.db.GetEmployeeByID(id); err != nil {
		return err
	}

	return s.db.DeleteEmployee(id)
}

// ListEmployee retrieves a list of all cloud providers
func (s *Service) ListEmployee() ([]*models.Employee, error) {
	employee, err := s.db.ListEmployee()
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// UpdateEmployee updates an existing employee with the provided information.
func (s *Service) UpdateEmployee(employee *models.Employee) (string, error) {
	employeeID, err := s.db.GetEmployeeByID(employee.ID)
	if err != nil {
		return "", err
	}

	if employee.ID != employeeID.ID {
		return "", domainerr.NewAPIError(domainerr.NotFound, "employee with ID does not exist")

	}

	existingEmployee, err := s.db.AddEmployee(employee)
	if err != nil {
		return "", err
	}

	return existingEmployee, nil
}
