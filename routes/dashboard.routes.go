package routes

import (
	"github.com/gin-gonic/gin"
	"glm-go-project/controllers"
)

type DashboardRouteController struct {
	dashboardController controllers.DashboardController
}

func NewDashboardRouteController(dashboardController controllers.DashboardController) DashboardRouteController {
	return DashboardRouteController{dashboardController}
}

func (uc *DashboardRouteController) DashboardRoute(rg *gin.RouterGroup) {

	router := rg.Group("dashboard")

	router.GET("/test", uc.dashboardController.GetDateDashboard)
}
