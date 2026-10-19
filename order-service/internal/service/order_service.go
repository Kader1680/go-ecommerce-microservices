package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/yourname/order-service/internal/models"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
	ErrOrderNotFound   = errors.New("order not found")
)

type OrderRepository interface {
	Save(order *models.Order) error
	FindByID(id string) (*models.Order, error)
	FindByCustomer(customerID string) ([]*models.Order, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(customerID, productID string, quantity int) (*models.Order, error) {
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	order := &models.Order{
		ID:         generateID(),
		CustomerID: customerID,
		ProductID:  productID,
		Quantity:   quantity,
		Status:     models.StatusPending,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Save(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetOrder(id string) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrOrderNotFound
	}
	return order, nil
}

func (s *OrderService) GetCustomerOrders(customerID string) ([]*models.Order, error) {
	return s.repo.FindByCustomer(customerID)
}

func (s *OrderService) CancelOrder(orderID string) error {
	order, err := s.repo.FindByID(orderID)
	if err != nil {
		return ErrOrderNotFound
	}
	if order.Status == models.StatusConfirmed {
		return errors.New("cannot cancel a confirmed order")
	}
	order.Status = models.StatusCancelled
	return s.repo.Save(order)
}

func generateID() string {
	return fmt.Sprintf("ord_%d", time.Now().UnixNano())
}
