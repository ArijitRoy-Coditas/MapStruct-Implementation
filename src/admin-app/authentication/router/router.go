package router

import (
	"authentication/business"
	"authentication/common/constants"
	"authentication/handlers"
	"authentication/middleware"
	"authentication/repositories"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	useDBMocks := false
	router := gin.New()
	router.Use(middleware.AuthMiddleware())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowMethods: []string{"POST", "GET"},
	}))

	createUserRepositories := repositories.NewGetCreateUserRepositories(useDBMocks)
	createUserService := business.NewGetCreateUserService(createUserRepositories)
	createUserController := handlers.NewGetCreateUserController(createUserService)

	router.POST(constants.CreateUser, createUserController.HandleCreateUser)
	return router
}
