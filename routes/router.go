package routes

import (
    "database/sql"
    "github.com/Bondrewdq/milktea-ordering-app/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
    router := gin.Default()

    // 初始化控制器
    teaCtrl := &controllers.TeaController{DB: db}

    // API 路由组
    api := router.Group("/api")
    {
        api.GET("/teas", teaCtrl.GetAllTeas)
        api.POST("/teas", teaCtrl.CreateTea)
    }

    return router
}