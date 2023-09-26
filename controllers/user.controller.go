package controllers

import (
	"github.com/gin-gonic/gin"
	"glm-go-project/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// @Summary		Получение пользователя по Id
// @Description	Получить пользователя по ИД
// @ID				get-user-from-id
// @Accept			json
// @Produce		json
// @Param id   path int true "Id"
// @Success		200		{object}	model.UserInfo
// @Router			/user/get-user-from-id/{id} [get]
func (uc *UserController) GetUserFromId(ctx *gin.Context) {
	userId := ctx.Param("id")

	var user model.UserInfo
	result := uc.DB.First(&user, "id = ?", userId)
	if result.Error != nil {

		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

// @Summary		Получение списка пользователей
// @Description	Получить пользователя по ИД
// @ID				get-users
// @Accept			json
// @Produce		json
// @Param   limit      query    int     true        "Limit"
// @Param   page      query    int     true        "Page"
// @Success		200		{array}	model.UserInfo
// @Router			/user/get-users/ [get]
func (pc *UserController) GetUsers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []model.UserInfo
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}
