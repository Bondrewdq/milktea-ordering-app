package main

import (
    "log"
    
    "github.com/Bondrewdq/milktea-ordering-app/config"
    "github.com/Bondrewdq/milktea-ordering-app/kafka"
    "github.com/Bondrewdq/milktea-ordering-app/services"
    "github.com/Bondrewdq/milktea-ordering-app/controllers"
    "github.com/Bondrewdq/milktea-ordering-app/routes"
)

func main() {
    // 1. 初始化数据库连接
    db := config.InitDB()
    
    // 2. 初始化Kafka生产者（必须在启动服务前调用）
    kafka.InitProducer()
    defer kafka.CloseProducer() // 确保程序退出时关闭连接

    // 3. 初始化Service层（注意：不再需要传递kafka实例）
    orderService := services.NewOrderService(db)
    // milkTeaService := services.NewMilkTeaService(db)
    
    // 4. 初始化Controller层
    orderCtrl := controllers.NewOrderController(orderService)
    milkTeaCtrl := &controllers.TeaController{DB: db} // 使用已有的控制器
    
    // 5. 设置路由
    router := routes.SetupRouter(orderCtrl, milkTeaCtrl)

    log.Println("Server starting on :8080...")
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}