package repository

import (
	"errors"
	"sync"

	"github.com/yourname/order-service/internal/models"
)

var ErrNotFound = errors.New("order not found")

// in-memory repo - can be swapped for Postgres later without touching the service
type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*models.Order),
	}
}

func (r *OrderRepository) Save(order *models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

func (r *OrderRepository) FindByID(id string) (*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.orders[id]
	if !ok {
		return nil, ErrNotFound
	}
	return order, nil
}

func (r *OrderRepository) FindByCustomer(customerID string) ([]*models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*models.Order
	for _, o := range r.orders {
		if o.CustomerID == customerID {
			result = append(result, o)
		}
	}
	return result, nil
}
