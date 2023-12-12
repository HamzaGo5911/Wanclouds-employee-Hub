package service

import "context"

// Close disconnects the mongo client
func (s *Service) Close(ctx context.Context) error {
	return s.db.Disconnect(ctx)
}
