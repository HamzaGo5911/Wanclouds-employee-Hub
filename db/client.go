package db

import (
	"context"
	"log"

	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	AddEmployee(employee *models.Employee) (string, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	DeleteEmployee(id string) error
	ListEmployee() ([]*models.Employee, error)
	AddOffice(office *models.Office) (string, error)
	GetOfficeByID(id string) (*models.Office, error)
	DeleteOffice(id string) error
	ListOffice() ([]*models.Office, error)
	AddHome(home *models.Home) (string, error)
	GetHomeByID(id string) (*models.Home, error)
	DeleteHome(id string) error
	ListHome() ([]*models.Home, error)
	Disconnect(ctx context.Context) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
