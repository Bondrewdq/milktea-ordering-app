package models

import "gorm.io/gorm"

type Milktea struct {
    gorm.Model  // 自动包含 ID, CreatedAt, UpdatedAt, DeletedAt

    Name      string    `gorm:"not null" json:"name"`
    Price     float64   `gorm:"not null" json:"price"`
}