package service

import (
	"github.com/kashifkhan0771/HostService/models"
)

// AddHost add into db.
func (s *Service) AddHost(host *models.Host) (string, error) {
	return s.db.AddHost(host)
}

// DeleteHost delete from db.
func (s *Service) DeleteHost(id string) error {
	if _, err := s.db.GetHost(id); err != nil {
		return err
	}

	return s.db.DeleteHost(id)
}

// RetrieveHost get host from db.
func (s *Service) RetrieveHost(id string) (*models.Host, error) {
	host, err := s.db.GetHost(id)
	if err != nil {
		return nil, err
	}

	return host, nil
}

// UpdateHost update in db.
func (s *Service) UpdateHost(host *models.Host, id string) error {
	if _, err := s.RetrieveHost(id); err != nil {
		return err
	}

	return s.db.UpdateHost(host)
}
