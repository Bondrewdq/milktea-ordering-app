package controllers

import (
    "net/http"
    "github.com/Bondrewdq/milktea-ordering-app/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type TeaController struct {
    DB *gorm.DB
}


// 获取所有奶茶
func (tc *TeaController) GetAllTeas(c *gin.Context) {
    var teas []models.Milktea
    if result := tc.DB.Find(&teas); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, teas)
}


// 添加奶茶
func (tc *TeaController) CreateTea(c *gin.Context) {
    var tea models.Milktea
    if err := c.ShouldBindJSON(&tea); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := tc.DB.Create(&tea); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, tea)
}


// 根据ID查询
func (tc *TeaController) GetTeaByID(c *gin.Context) {
    id := c.Param("id")
    var tea models.Milktea
    if result := tc.DB.First(&tea, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tea not found"})
        return
    }
    c.JSON(http.StatusOK, tea)
}


// 更新奶茶
func (tc *TeaController) UpdateTea(c *gin.Context) {
    id := c.Param("id")
    var tea models.Milktea
    if err := c.ShouldBindJSON(&tea); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := tc.DB.Model(&models.Milktea{}).Where("id = ?", id).Updates(tea); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}


// 删除奶茶
func (tc *TeaController) DeleteTea(c *gin.Context) {
    id := c.Param("id")
    if result := tc.DB.Delete(&models.Milktea{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}