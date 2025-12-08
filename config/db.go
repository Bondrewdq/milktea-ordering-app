package config

import (
    "log"
    "github.com/Bondrewdq/milktea-ordering-app/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("milktea.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // 自动迁移表结构（创建/更新表）
    db.AutoMigrate(&models.Milktea{})
    //db.AutoMigrate(&models.Order{}) // 示例：后续可扩展订单模型

    log.Println("Database initialized with GORM")
    return db
}