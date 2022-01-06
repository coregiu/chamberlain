package notebook

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
	"time"
)

type SummaryBook struct {
	BookId       string    `gorm:"column:BOOK_ID"`
	ParentBookId string    `gorm:"column:PARENT_BOOK_ID"`
	Username     string    `gorm:"column:USERNAME"`
	Content      string    `gorm:"column:CONTENT"`
	BookTime     time.Time `gorm:"column:BOOK_TIME"`
}

type SummaryBookMgmt interface {
	AddSummaryBook() error
	BatchAddSummaryBook(summaryBook *[]SummaryBook) error
	UpdateSummaryBook() error
	DeleteSummaryBook() error
	/*Get details*/
	GetSummaryBooks() ([]SummaryBook, error)
	GetSummaryBookContent() error
}

func (SummaryBook) TableName() string {
	return "SUMMARY_BOOKS"
}

func (summaryBook *SummaryBook) AddSummaryBook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&summaryBook)
	return result.Error
}

func (summaryBook SummaryBook) BatchAddSummaryBook(summaryBooks *[]SummaryBook) error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.CreateInBatches(summaryBook, len(*summaryBooks))
	return result.Error
}

func (summaryBook *SummaryBook) UpdateSummaryBook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	dataSet := db.Model(&SummaryBook{}).Where("BOOK_ID = ? AND USERNAME=?", summaryBook.BookId, summaryBook.Username)
	if summaryBook.Content != "" {
		result := dataSet.Update("CONTENT", summaryBook.Content)
		if result.Error != nil {
			return result.Error
		}
	}
	if summaryBook.ParentBookId != "" {
		result := dataSet.Update("PARENT_BOOK_ID", summaryBook.ParentBookId)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (summaryBook *SummaryBook) DeleteSummaryBook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&SummaryBook{}, "BOOK_ID = ? AND USERNAME=?", summaryBook.BookId, summaryBook.Username)
	return result.Error
}

func (summaryBook *SummaryBook) GetSummaryBooks() ([]SummaryBook, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	summaryBooks := make([]SummaryBook, 0)
	result := db.Select("BOOK_ID, PARENT_BOOK_ID, USERNAME, BOOK_TIME").Where("USERNAME=?", summaryBook.Username).Or("PARENT_BOOK_ID, BOOK_TIME DESC").Find(&summaryBook)
	return summaryBooks, result.Error
}

func (summaryBook *SummaryBook) GetSummaryBooksContent() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Where("BOOK_ID=? AND USERNAME=?", summaryBook.BookId, summaryBook.Username).Find(&summaryBook)
	return result.Error
}
