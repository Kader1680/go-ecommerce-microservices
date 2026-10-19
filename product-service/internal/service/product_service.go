package service

import (
	"errors"

	"github.com/yourname/product-service/internal/models"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

type ProductRepository interface {
	Save(p *models.Product) error
	FindByID(id string) (*models.Product, error)
	FindAll() ([]*models.Product, error)
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(id string) (*models.Product, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrProductNotFound
	}
	return p, nil
}

func (s *ProductService) ListProducts() ([]*models.Product, error) {
	return s.repo.FindAll()
}

// ReserveStock is what Order Service will call (via HTTP) before confirming an order.
func (s *ProductService) ReserveStock(id string, quantity int) error {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return ErrProductNotFound
	}
	if p.Stock < quantity {
		return ErrInsufficientStock
	}
	p.Stock -= quantity
	return s.repo.Save(p)
}
