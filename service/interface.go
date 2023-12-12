package service

import "github.com/HamzaGo5911/wanclouds-employee-hub/db"

// Service initializes our database instance
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
