package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Bondrewdq/milktea-ordering-app/controllers/request"
	"github.com/Bondrewdq/milktea-ordering-app/controllers/response"
	"github.com/Bondrewdq/milktea-ordering-app/services"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(svc services.OrderService) *OrderController {
	return &OrderController{orderService: svc}
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := ctrl.orderService.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.OrderResponse{
		OrderID:    strconv.FormatUint(uint64(order.ID), 10),
		Status:     strconv.Itoa(order.Status),
		TotalPrice: order.TotalPrice,
		Message:    "订单创建成功，等待支付",
		CreatedAt:  order.CreatedAt,
	})
}