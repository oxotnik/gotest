package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type DashboardController struct {
	DB *gorm.DB
}

func NewDashboardController(DB *gorm.DB) DashboardController {
	return DashboardController{DB}
}

// @Summary		Метод тест - не реализован
// @Description	не реализован
// @ID				dashboard-test
// @Accept			json
// @Produce		json
// @Success		200		{string}	string			"ok"
// @Router			/dashboard/test [get]
func (dc *DashboardController) GetDateDashboard(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Не реализован"})
}
