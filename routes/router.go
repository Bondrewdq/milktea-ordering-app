package routes

import (
    "github.com/Bondrewdq/milktea-ordering-app/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(orderCtrl *controllers.OrderController, teaCtrl *controllers.TeaController) *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        // 奶茶相关接口
        api.GET("/teas", teaCtrl.GetAllTeas)
        api.GET("/teas/:id", teaCtrl.GetTeaByID)
        api.POST("/teas", teaCtrl.CreateTea)
        api.PUT("/teas/:id", teaCtrl.UpdateTea)
        api.DELETE("/teas/:id", teaCtrl.DeleteTea)

        // 订单相关接口
        api.POST("/orders", orderCtrl.CreateOrder)
    }

    return router
}