package server

import (
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin"
)

const AuthHeader string = "X-AUTH-TOKEN"

func StartServer() {
	log.Info("Begin to start web server...")
	router := gin.Default()
	apiRoute(router)
	router.Run()
	log.Info("Finished web server start...")
}

func apiRoute(router *gin.Engine) {
	userIndex := router.Group("/users")
	{
		userIndex.POST("", AddUserHandler())
		userIndex.PUT("", AuthHandler(), UpdateUserHandler())
		userIndex.DELETE("", AuthHandler(), DeleteUserHandler())
		userIndex.GET("", AuthHandler(), GetUsersHandler())
		userIndex.GET("/count", AuthHandler(), GetUsersCountHandler())
		userIndex.GET("/user/:username", AuthHandler(), GetUserHandler())
		userIndex.POST("/login", LoginHandler())
		userIndex.POST("/logout", AuthHandler(), LogoutHandler())
	}

	inputIndex := router.Group("/inputs")
	{
		inputIndex.POST("", AuthHandler(), AddInputHandler())
		inputIndex.PUT("", AuthHandler(), UpdateInputHandler())
		inputIndex.DELETE("", AuthHandler(), DeleteInputHandler())
		inputIndex.GET("", AuthHandler(), GetInputsHandler())
		inputIndex.GET("/count", AuthHandler(), GetInputsCountHandler())
		inputIndex.GET("/statistic", AuthHandler(), GetStatisticHandler())
		inputIndex.GET("/statistic/month", AuthHandler(), GetStatisticByMonthHandler())
		inputIndex.GET("/statistic/type", AuthHandler(), GetStatisticByTypeHandler())
	}
}