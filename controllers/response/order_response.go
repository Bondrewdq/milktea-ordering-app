package response

import "time"

type OrderResponse struct {
	OrderID    string    `json:"order_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}