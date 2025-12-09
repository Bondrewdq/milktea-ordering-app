package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Bondrewdq/milktea-ordering-app/controllers/request"
	"github.com/Bondrewdq/milktea-ordering-app/kafka"
	"github.com/Bondrewdq/milktea-ordering-app/models"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req request.CreateOrderRequest) (*models.Order, error)
}

type orderService struct {
	db      *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{db: db}
}

func (s *orderService) CreateOrder(ctx context.Context, req request.CreateOrderRequest) (*models.Order, error) {
	var order models.Order
	
	// 使用事务保证一致性
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 计算总价并扣库存
		for _, item := range req.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				return err
			}

			// 创建订单
			order = models.Order{
				UserID:     req.UserID,
				ProductID:  item.ProductID,
				Quantity:   item.Quantity,
				TotalPrice: product.Price * float64(item.Quantity),
				Status:     0,
			}
			
			// 扣减库存
			if err := tx.Model(&product).Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return err
			}
			
			if err := tx.Create(&order).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 异步发送Kafka消息（削峰）
	go s.sendOrderEvent(order.ID)

	return &order, nil
}

// 发送订单创建事件
func (s *orderService) sendOrderEvent(orderID uint) {
	msg, _ := json.Marshal(map[string]interface{}{
		"event":    "order.created",
		"order_id": orderID,
		"timestamp": time.Now().Unix(),
	})
	kafka.SendOrderMessage(orderID, string(msg))
}