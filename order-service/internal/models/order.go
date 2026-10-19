package models

import "time"

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	ProductID  string      `json:"product_id"`
	Quantity   int         `json:"quantity"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
}
