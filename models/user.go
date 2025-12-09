package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	
	Username string `json:"username" gorm:"unique;not null"` // 用户名，唯一
	Password string `json:"password" gorm:"not null"`        // 密码
	Phone    string `json:"phone"`                           // 手机号

	// 关联关系：一个用户有多个订单（可选，Gorm会自动处理，不写也行，写了查询方便）
	Orders   []Order `json:"orders,omitempty" gorm:"foreignKey:UserID"` 
}