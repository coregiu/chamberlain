package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const AUTH_HEADER string = "X-AUTH-TOKEN"

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Info("begin to check auth")
		tokenId := context.GetHeader(AUTH_HEADER)
		if tokenId == "" {
			context.JSON(400, gin.H{
				"result": "no authorization for no auth token",
			})
			context.Abort()
			return
		}
		token := auth.Token{}
		token.TokenId = tokenId
		log.Info("request url = %s", context.Request.RequestURI)
		url := context.Request.RequestURI
		arr := strings.Split(url, "/")
		log.Info("operation = %s", arr[1])
		isAuthed, err := token.CheckAuth(arr[1])
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

func GetUserHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Info("get user information by username")
		username := context.Param("username")
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
	}
}

func GetUsersHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit, _ := context.GetQuery("limit")
		offset, _ := context.GetQuery("offset")
		limitInt, _ := strconv.Atoi(limit)
		offsetInt, _ := strconv.Atoi(offset)

		users := make([]auth.User, 0)
		user := new(auth.User)
		users, err := user.GetUsers(offsetInt, limitInt)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, users)
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
			context.JSON(200, gin.H{
				"count": usersCount,
			})
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
		user.Role = "User"
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
		tokenId := context.GetHeader(AUTH_HEADER)
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

func AddInputHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "inputs successfully",
		})
	}
}
