package server

import (
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin"
)

const AuthHeader string = "X-AUTH-TOKEN"

func StartServer() {
	log.Info("Begin to start web server...")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	apiRoute(router)
	_ = router.Run()
	log.Info("Finished web server start...")
}

func apiRoute(router *gin.Engine) {
	userIndex := router.Group("/users", AuthHandler())
	{
		userIndex.POST("", AddUserHandler())
		userIndex.PUT("", UpdateUserHandler())
		userIndex.DELETE("", DeleteUserHandler())
		userIndex.GET("", GetUsersHandler())
		userIndex.GET("/count", GetUsersCountHandler())
		userIndex.GET("/token", GetUserByTokenHandler())
		userIndex.PUT("/password", RestPasswordHandler())
		userIndex.POST("/login", LoginHandler())
		userIndex.POST("/logout", LogoutHandler())
	}
	inputIndex := router.Group("/inputs", AuthHandler())
	{
		inputIndex.POST("", AddInputHandler())
		inputIndex.PUT("", UpdateInputHandler())
		inputIndex.DELETE("", DeleteInputHandler())
		inputIndex.GET("", GetInputsHandler())
		inputIndex.GET("/count", GetInputsCountHandler())
		inputIndex.GET("/statistic", GetStatisticHandler())
		inputIndex.GET("/statistic/month", GetStatisticByMonthHandler())
		inputIndex.GET("/statistic/type", GetStatisticByTypeHandler())
	}
	blogsIndex := router.Group("/blogs", AuthHandler())
	{
		blogsIndex.POST("", BlogsHandler())
	}
}
