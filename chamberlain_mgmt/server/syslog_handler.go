package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
			mapToken, err := token.GetTokenById()
			if err != nil {
				username = context.Writer.Header().Get("USERNAME")
			} else {
				username = mapToken.User.Username
			}
		} else {
			username = context.Writer.Header().Get("USERNAME")
		}

		syslog.OpTime = time.Now()
		syslog.LogId = syslog.OpTime.Unix()
		syslog.Operation = context.Request.Method+":"+context.Request.URL.Path
		syslog.OpClient = context.Request.Host
		syslog.OpResult = strconv.Itoa(status)
		syslog.Description = description
		syslog.Username = username
		err := syslog.AddSyslog()
		if err != nil {
			log.Error("failed to record log.")
		}
	}
}
