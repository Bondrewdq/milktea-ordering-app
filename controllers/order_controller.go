// Package controllers 提供了应用程序的HTTP控制器层
// 包含订单、用户、产品等业务逻辑的HTTP接口处理
package controllers

import (
	"net/http"

	"github.com/Bondrewdq/milktea-ordering-app/controllers/request"
	// "github.com/Bondrewdq/milktea-ordering-app/controllers/response"
	"github.com/Bondrewdq/milktea-ordering-app/services"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// CreateOrder 下单接口
// @Summary 创建订单
// @Description 用户下单购买奶茶
// @Tags orders
// @Accept json
// @Produce json
// @Param request body request.CreateOrderRequest true "下单请求"
// @Success 200 {object} response.CreateOrderResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := ctrl.service.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": order.ID,
		"message":  "Order created successfully",
	})
}