package service

import (
	"e-comerce/internal/adapters/repositories"
	"e-comerce/internal/models"
	"log"
	"strings"
)

type store struct {
	storage repositories.StorageGetter
}

func NewServiceGetter(storage repositories.StorageGetter) ServiceGetter {
	return &store{
		storage: storage,
	}
}

func (s *store) GetOneByName(productName string) (*models.Product, error) {
	productName = strings.ToLower(productName)
	product, err := s.storage.GetOneByName(productName)
	if err != nil {
		log.Printf("Error in GetOne method in service: %v\n", err)
		return nil, err
	}
	return product, nil
}

func (s *store) GetAll() ([]*models.Product, error) {
	products, err := s.storage.GetAll()
	if err != nil {
		log.Printf("error in get all service: %v\n", err)
		return nil, err
	}
	return products, nil
}
