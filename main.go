package main

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"glm-go-project/controllers"
	_ "glm-go-project/docs"
	docs "glm-go-project/docs"
	"glm-go-project/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"glm-go-project/initializers"
)

var (
	server                   *gin.Engine
	DashboardController      controllers.DashboardController
	DashboardRouteController routes.DashboardRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	DashboardController = controllers.NewDashboardController(initializers.DB)
	DashboardRouteController = routes.NewDashboardRouteController(DashboardController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

// @title           Go First Microservice test
// @version         1.0
// @description     Реализовал простой микросервис,  посмотрел возможности gin
// @description     Посмотрел orm gome,  подключил postgres, т.к. бд у меня была, то по подходу db first, сформировал модель в коде и заюзал простые методы работы с данными.
// @description     Добавил swagger
// @description     Настроил окружение на сервере и развернул тестовый микросервис.

// @contact.name   API Support
// @contact.email  vmatsiukhin@gmail.com

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	docs.SwaggerInfo.BasePath = "/api/v1"

	router := server.Group("/api/v1")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	//v1 := server.Group("/api/v1")
	//{
	//	eg := v1.Group("/example")
	//	{
	//		eg.GET("/helloworld", nil)
	//	}
	//}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	DashboardRouteController.DashboardRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

}
