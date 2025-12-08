package main

import (
    "log"
    "github.com/Bondrewdq/milktea-ordering-app/config"
    "github.com/Bondrewdq/milktea-ordering-app/routes"
)

func main() {
    // 初始化数据库
    db := config.InitDB()
    defer db.Close()

    // 初始化 Gin 路由
    router := routes.SetupRouter(db)

    // 启动服务器
    log.Println("Server starting on :8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}