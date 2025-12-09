package models

import "gorm.io/gorm"

// Order 模型代表了一次购买交易记录。
type Order struct {
	gorm.Model

	UserID 		uint	 `json:"user_id" gorm:"not null;index"` 		// 外键关联：谁下的单？
	// Gorm 关联定义，用于预加载 (Preload) 查询用户数据
	User   		User 	`json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` 

	ProductID 	uint    `json:"product_id" gorm:"not null;index"`		// 外键关联：买了哪个商品？
	Product   	Product `json:"product" gorm:"foreignKey:ProductID"`	// Gorm 关联定义，用于查询商品详情
	
	Quantity    int     `json:"quantity" gorm:"not null"`     			// 购买数量
	TotalPrice  float64 `json:"total_price"`                  			// 总价 (Quantity * Price)
	
	// 核心状态流转字段：Kafka 消费者会监听此状态
	// 0 - 已下单 / 待制作 (Producer 初始设置)
	// 1 - 制作中 (Consumer 接收消息后设置)
	// 2 - 制作完成 / 待取餐 (Consumer 模拟耗时后设置)
	// 3 - 已取消
	Status int `json:"status" gorm:"default:0"` 
}