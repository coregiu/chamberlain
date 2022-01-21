package server

import (
	"chamberlain_mgmt/auth"
	"chamberlain_mgmt/log"
	note "chamberlain_mgmt/notebook"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddSummaryBookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		summaryBook := &note.SummaryBook{}
		summaryBooks := make([]note.SummaryBook, 0)
		err := context.BindJSON(&summaryBooks)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("summaryBook length =%s", fmt.Sprint(len(summaryBooks)))
		username := getLoginUsername(context)

		for index := range summaryBooks {
			summaryBooks[index].Username = username
		}

		err = summaryBook.BatchAddSummaryBook(&summaryBooks)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Add summaryBook successfully.",
			})
		}
	}
}

func UpdateSummaryBookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		summaryBook := note.SummaryBook{}
		err := context.BindJSON(&summaryBook)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("InputTime =%s", fmt.Sprint(summaryBook.BookId))
		username := getLoginUsername(context)
		summaryBook.Username = username

		err = summaryBook.UpdateSummaryBook()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Update summaryBook successfully.",
			})
		}
	}
}

func DeleteSummaryBookHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		noteSummary := note.SummaryBook{}
		err := context.BindJSON(&noteSummary)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("BookId =%s", fmt.Sprint(noteSummary.BookId))
		username := getLoginUsername(context)
		noteSummary.Username = username
		err = noteSummary.DeleteSummaryBookWithChildren()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Delete noteSummary successfully.",
			})
		}
	}
}

func GetSummaryBooksHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := getIntParam(context, "limit", 10)
		offset := getIntParam(context, "offset", 0)
		queryFinishTime, _ := context.GetQuery("finish_time")
		status, _ := context.GetQuery("status")
		log.Info("finish time = %s, status = %s, limit = %d, offset = %d", queryFinishTime, status, limit, offset)

		username := getLoginUsername(context)

		notebook := note.SummaryBook{}
		notebook.Username = username
		notebooks, err := notebook.GetSummaryBooks()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, notebooks)
		}
	}
}

func GetSummaryBookContentHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		bookId, _ := context.GetQuery("book_id")
		log.Info("book_id = %s", bookId)

		username := getLoginUsername(context)

		summaryBook := note.SummaryBook{}
		summaryBook.Username = username
		summaryBook.BookId = bookId
		err := summaryBook.GetSummaryBooksContent()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, summaryBook)
		}
	}
}

func getLoginUsername(context *gin.Context) string {
	username := ""
	tokenId := context.Request.Header.Get("X-AUTH-TOKEN")
	token := auth.Token{}
	token.TokenId = tokenId
	mapToken, err := token.GetTokenById()
	if err == nil {
		username = mapToken.User.Username
	}
	return username
}
