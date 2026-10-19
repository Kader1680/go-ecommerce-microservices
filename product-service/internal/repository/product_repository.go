package repository

import (
	"errors"
	"sync"

	"github.com/yourname/product-service/internal/models"
)

var ErrNotFound = errors.New("product not found")

type ProductRepository struct {
	mu       sync.RWMutex
	products map[string]*models.Product
}

func NewProductRepository() *ProductRepository {
	repo := &ProductRepository{products: make(map[string]*models.Product)}
	// seed data so Order Service has something real to call
	repo.products["p1"] = &models.Product{ID: "p1", Name: "Laptop", Price: 1200.0, Stock: 10}
	repo.products["p2"] = &models.Product{ID: "p2", Name: "Mouse", Price: 25.0, Stock: 100}
	return repo
}

func (r *ProductRepository) Save(p *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[p.ID] = p
	return nil
}

func (r *ProductRepository) FindByID(id string) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.products[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (r *ProductRepository) FindAll() ([]*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*models.Product
	for _, p := range r.products {
		result = append(result, p)
	}
	return result, nil
}
