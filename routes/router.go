package routes

import (
    "github.com/Bondrewdq/milktea-ordering-app/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    teaCtrl := &controllers.TeaController{DB: db}

    api := router.Group("/api")
    {
        api.GET("/teas", teaCtrl.GetAllTeas)
        api.GET("/teas/:id", teaCtrl.GetTeaByID)
        api.POST("/teas", teaCtrl.CreateTea)
        api.PUT("/teas/:id", teaCtrl.UpdateTea)
        api.DELETE("/teas/:id", teaCtrl.DeleteTea)
    }

    return router
}