package service

import (
	"e-comerce/internal/adapters/repositories"
	"e-comerce/internal/models"
	"e-comerce/pkg"
	"errors"
	"log"
	"strings"
)

type editStore struct {
	storage repositories.StorageEditor
}

func NewServiceEditor(storage repositories.StorageEditor) ServiceEditor {
	return &editStore{
		storage: storage,
	}
}

// CreateOne implements ServiceEditor
func (s *editStore) CreateOne(p *models.Product) (int, error) {
	if !pkg.IsValid(p.Price) {
		return 0, errors.New("Price is not valid")
	}
	p.Name = strings.ToLower(p.Name)
	id, err := s.storage.CreateOne(p)
	if err != nil {
		log.Printf("error in create one method service: %v\n", err)
		return 0, err
	}
	return id, err
}

// DeleteOne implements ServiceEditor
func (s *editStore) DeleteOne(id int) error {
	if !pkg.IsValid(id) {
		log.Printf("id for delete product is not valid")
		return errors.New("id for delete product is not valid")
	}
	if err := s.storage.DeleteOne(id); err != nil {
		log.Printf("error in delete one method service: %v\n", err)
		return err
	}
	return nil
}

// UpdateOne implements ServiceEditor
func (s *editStore) UpdateOne(p *models.Product) error {
	if !pkg.IsValid(p.Price) {
		return errors.New("Price is not valid")
	}
	p.Name = strings.ToLower(p.Name)
	if err := s.storage.UpdateOne(p); err != nil {
		log.Printf("error in update one method: %v\n", err)
		return err
	}
	return nil
}
