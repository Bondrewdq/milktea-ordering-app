package request

type CreateOrderRequest struct {
	UserID uint               `json:"user_id" binding:"required"`
	Items  []OrderItemRequest `json:"items" binding:"required,min=1"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}