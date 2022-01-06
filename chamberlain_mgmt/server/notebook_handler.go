package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	note "chamberlain_mgmt/notebook"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddNotebookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		notebook := &note.Notebook{}
		notebooks := make([]note.Notebook, 0)
		err := context.BindJSON(&notebooks)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("notebook length =%s", fmt.Sprint(len(notebooks)))
		tokenId := context.Request.Header.Get("X-AUTH-TOKEN")
		token := auth.Token{}
		token.TokenId = tokenId
		mapToken, err := token.GetToken()
		if err == nil {
			username := mapToken.User.Username
			for index := range notebooks {
				notebooks[index].Username = username
			}
		}
		err = notebook.BatchAddNotebook(&notebooks)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Add notebook successfully.",
			})
		}
	}
}

func UpdateNotebookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		notebook := note.Notebook{}
		err := context.BindJSON(&notebook)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("InputTime =%s", fmt.Sprint(notebook.NoteId))
		err = notebook.UpdateNotebook()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Update notebook successfully.",
			})
		}
	}
}

func DeleteNotebookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		notebook := note.Notebook{}
		err := context.BindJSON(&notebook)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("NoteId =%s", fmt.Sprint(notebook.NoteId))
		err = notebook.DeleteNotebook()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Delete notebook successfully.",
			})
		}
	}
}

func GetNotebooksHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := getIntParam(context, "limit", 10)
		offset := getIntParam(context, "offset", 0)
		queryFinishTime, _ := context.GetQuery("finish_time")
		status, _ := context.GetQuery("status")
		log.Info("finish time = %s, status = %s, limit = %d, offset = %d", queryFinishTime, status, limit, offset)

		username := ""
		tokenId := context.Request.Header.Get("X-AUTH-TOKEN")
		token := auth.Token{}
		token.TokenId = tokenId
		mapToken, err := token.GetToken()
		if err == nil {
			username = mapToken.User.Username
		}

		notebook := note.Notebook{}
		notebook.Username = username
		notebook.Status = status
		notebooks, err := notebook.GetNotebooks(queryFinishTime, limit, offset)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, notebooks)
		}
	}
}
