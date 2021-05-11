package server

import (
	"chamberlain_mgmt/blog"
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin"
)

func BlogsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		blogWorkPath, blogRepos := config.GetBlogsConfig()
		blogs := new(blog.Blogs)
		blogs.WorkPath = blogWorkPath
		blogs.BlogRepos = blogRepos

		err := blogs.CleanWorkSpace()
		if err != nil {
			log.Error(err.Error())
			context.JSON(500, err.Error())
		}
		err = blogs.HandleBlogs()
		if err != nil {
			log.Error(err.Error())
			context.JSON(500, err.Error())
		}
		context.JSON(200, gin.H{
			"result": "logout successfully",
		})
	}
}
