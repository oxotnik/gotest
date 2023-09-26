package routes

import (
	"github.com/gin-gonic/gin"
	"glm-go-project/controllers"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("user")

	router.GET("/get-user-from-id/:id", uc.userController.GetUserFromId)
	router.GET("/get-users", uc.userController.GetUsers)
}
