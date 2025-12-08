package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Bondrewdq/milktea-ordering-app/models"

	"github.com/gin-gonic/gin"
)

type TeaController struct {
    DB *sql.DB
}

// 获取所有奶茶
func (tc *TeaController) GetAllTeas(c *gin.Context) {
    rows, err := tc.DB.Query("SELECT id, name, price, created_at FROM teas")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var teas []models.Milktea
    for rows.Next() {
        var tea models.Milktea
        if err := rows.Scan(&tea.ID, &tea.Name, &tea.Price, &tea.CreatedAt); err != nil {
            continue
        }
        teas = append(teas, tea)
    }

    c.JSON(http.StatusOK, teas)
}

// 添加奶茶（示例）
func (tc *TeaController) CreateTea(c *gin.Context) {
    var tea models.Milktea
    if err := c.ShouldBindJSON(&tea); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := tc.DB.Exec("INSERT INTO teas (name, price) VALUES (?, ?)", tea.Name, tea.Price)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    id, _ := result.LastInsertId()
    tea.ID = int(id)
    c.JSON(http.StatusCreated, tea)
}