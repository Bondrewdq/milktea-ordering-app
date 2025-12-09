package models

import "gorm.io/gorm"

// Product 模型代表了奶茶店提供的商品菜单。
type Product struct {
	gorm.Model        // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt
	
	Name        string  `json:"name" gorm:"not null"`                // 奶茶名，如 "珍珠奶茶"
	Price       float64 `json:"price" gorm:"not null"`               // 价格
	Description string  `json:"description"`                         // 描述
	// 关键字段：库存，防止超卖需要进行事务操作和并发控制
	Stock       int     `json:"stock" gorm:"not null;check:stock >= 0"` 
	ImageURL    string  `json:"image_url"`                           // 图片链接
}