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
		userIndex.POST("", AddUserHandler(), RecordLogHandler())
		userIndex.PUT("", UpdateUserHandler(), RecordLogHandler())
		userIndex.DELETE("", DeleteUserHandler(), RecordLogHandler())
		userIndex.GET("", GetUsersHandler())
		userIndex.GET("/count", GetUsersCountHandler())
		userIndex.GET("/token", GetUserByTokenHandler())
		userIndex.PUT("/password", RestPasswordHandler(), RecordLogHandler())
		userIndex.POST("/login", LoginHandler(), RecordLogHandler())
		userIndex.POST("/logout", LogoutHandler(), RecordLogHandler())
	}

	inputIndex := router.Group("/inputs", AuthHandler())
	{
		inputIndex.POST("", AddInputHandler(), RecordLogHandler())
		inputIndex.PUT("", UpdateInputHandler(), RecordLogHandler())
		inputIndex.DELETE("", DeleteInputHandler(), RecordLogHandler())
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

	syslogIndex := router.Group("/syslogs", AuthHandler())
	{
		syslogIndex.GET("", QuerySyslogHandler())
		syslogIndex.DELETE("", DeleteSyslogHandler(), RecordLogHandler())
	}

	notebookIndex := router.Group("/notebooks", AuthHandler())
	{
		notebookIndex.GET("", GetNotebooksHandler())
		notebookIndex.DELETE("", DeleteNotebookHandler(), RecordLogHandler())
		notebookIndex.POST("", AddNotebookHandler(), RecordLogHandler())
		notebookIndex.PUT("", UpdateNotebookHandler(), RecordLogHandler())
	}

	summaryBookIdx := router.Group("/summarybooks", AuthHandler())
	{
		summaryBookIdx.GET("", GetSummaryBooksHandler())
		summaryBookIdx.GET("/content", GetSummaryBookContentHandler())
		summaryBookIdx.DELETE("", DeleteSummaryBookHandler(), RecordLogHandler())
		summaryBookIdx.POST("", AddSummaryBookHandler(), RecordLogHandler())
		summaryBookIdx.PUT("", UpdateSummaryBookHandler(), RecordLogHandler())
	}
}
