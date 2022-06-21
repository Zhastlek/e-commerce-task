package service

import (
	"e-comerce/internal/models"
)

type ServiceGetter interface {
	GetAll() ([]*models.Product, error)
	GetOneByName(productName string) (*models.Product, error)
	GetOneById(productId string) (*models.Product, error)
}

type ServiceEditor interface {
	CreateOne(p *models.Product) (int, error)
	UpdateOne(p *models.Product) error
	DeleteOne(id int) error
}
