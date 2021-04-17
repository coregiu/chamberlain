package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Info("begin to check auth")
		tokenId := context.GetHeader(AuthHeader)

		token := auth.Token{}
		token.TokenId = tokenId
		log.Info("request url = %s", context.Request.RequestURI)
		url := context.Request.RequestURI
		indexOfParam := strings.Index(url, "?")
		if indexOfParam > 0 {
			url = url[0:indexOfParam]
		}
		log.Info("url = %s", url)
		isAuthed, err := token.CheckAuth(url)
		if err != nil || !isAuthed {
			context.JSON(400, gin.H{
				"result": "no authorization for expired token or has no permission",
			})
			context.Abort()
		} else {
			log.Info("authorization is ok")
			context.Next()
		}
	}
}

func GetUsersHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Info("get user information")
		username, _ := context.GetQuery("username")
		if username != "" {
			log.Info("username = " + username)
			user := new(auth.User)
			user.Username = username
			err := user.GetUser()
			if err != nil {
				log.Error("Failed to get the user %s", username)
				context.String(500, err.Error())
			} else {
				context.JSON(200, user)
			}
			context.Done()
		} else {

			limit, _ := context.GetQuery("limit")
			offset, _ := context.GetQuery("offset")
			limitInt, _ := strconv.Atoi(limit)
			offsetInt, _ := strconv.Atoi(offset)

			user := new(auth.User)
			users, err := user.GetUsers(offsetInt, limitInt)
			if err != nil {
				context.String(500, err.Error())
			} else {
				context.JSON(200, users)
			}
		}
	}
}

func GetUsersCountHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := new(auth.User)
		usersCount, err := user.GetUsersCount()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{"count": usersCount})
		}
	}
}

func AddUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := auth.User{}
		err := context.BindJSON(&user)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Info("username = " + user.Username)
		user.Role = "user"
		err = user.Adduser()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Add user successfully.",
			})
		}
	}
}

func UpdateUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := auth.User{}

		err := context.BindJSON(&user)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		err = user.UpdateUser()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Update user successfully.",
			})
		}
	}
}

func DeleteUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := auth.User{}

		err := context.BindJSON(&user)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Info("delete the user %s", user.Username)
		err = user.DeleteUser()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Delete user successfully.",
			})
		}
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := auth.User{}
		err := context.BindJSON(&user)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		result, err := user.CheckAuth()
		if err != nil || result == false {
			context.String(500, "failed to login")
		} else {
			token := auth.Token{}
			err := token.CreateNewToken(&user)
			if err != nil {
				log.Error("failed to create token")
				context.String(500, "failed to login")
			} else {
				context.JSON(200, token)
			}
		}
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenId := context.GetHeader(AuthHeader)
		if tokenId == "" {
			context.JSON(500, gin.H{
				"result": "no need logout",
			})
			return
		}
		token := auth.Token{}
		token.TokenId = tokenId
		token.DeleteToken()
		context.JSON(200, gin.H{
			"result": "logout successfully",
		})
	}
}
