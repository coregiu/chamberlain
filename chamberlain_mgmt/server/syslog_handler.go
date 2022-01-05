package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func QuerySyslogHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, _ := context.GetQuery("username")
		operation, _ := context.GetQuery("operation")
		limit, _ := context.GetQuery("limit")
		offset, _ := context.GetQuery("offset")
		limitInt, _ := strconv.Atoi(limit)
		offsetInt, _ := strconv.Atoi(offset)

		syslog := new(log.SysLog)
		syslogs, err := syslog.GetSyslog(username, operation, limitInt, offsetInt)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, syslogs)
		}
	}
}

func DeleteSyslogHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		syslog := &log.SysLog{}
		syslogs := make([]log.SysLog, 0)
		err := context.BindJSON(&syslogs)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("syslog length =%s", fmt.Sprint(len(syslogs)))
		err = syslog.DeleteSyslog(&syslogs)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Delete syslogs successfully.",
			})
		}
	}
}

func RecordLogHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		status := context.Writer.Status()
		log.Info("status = %d", status)
		log.Info("path = %s", context.Request.URL.Path)
		tokenId := context.Request.Header.Get("X-AUTH-TOKEN")
		log.Info("token = %s", tokenId)
		syslog := log.SysLog{}
		description := "success"
		if status >= 200 && status < 400 {
			description = "Success"
		} else if status == 401 {
			description = "Unauthorized"
		} else {
			description = "System error"
		}

		var username string
		if tokenId != "" {
			token := auth.Token{}
			token.TokenId = tokenId
			mapToken, err := token.GetToken()
			if err != nil {
				username = context.Writer.Header().Get("USERNAME")
			} else {
				username = mapToken.User.Username
			}
		} else {
			username = context.Writer.Header().Get("USERNAME")
		}
		err := syslog.RecordOperation(username, context.Request.Method+":"+context.Request.URL.Path, strconv.Itoa(status), description)
		if err != nil {
			log.Error("failed to record log.")
		}
	}
}
