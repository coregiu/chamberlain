package notebook

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Notebook struct {
	NoteId         string    `gorm:"column:NOTE_ID"`
	Username       string    `gorm:"column:USERNAME"`
	Content        string    `gorm:"column:CONTENT"`
	Level          string    `gorm:"column:LEVEL"`
	Owner          string    `gorm:"column:OWNER"`
	NoteTime       time.Time `gorm:"column:NOTE_TIME"`
	FinishTime     time.Time `gorm:"column:FINISH_TIME"`
	RealFinishTime time.Time `gorm:"column:REAL_FINISH_TIME"`
	Status         string    `gorm:"column:STATUS"`
}

type NotebookMgmt interface {
	AddNotebook() error
	BatchAddNotebook(notebooks *[]Notebook) error
	UpdateNotebook() error
	DeleteNotebook() error
	/*Get details*/
	GetNotebooks(finishTime string, limit int, offset int) ([]Notebook, error)
	GetNotebooksCount(finishTime time.Time, status string) (int64, error)
}

func (Notebook) TableName() string {
	return "NOTEBOOKS"
}

func (notebook *Notebook) AddNotebook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&notebook)
	return result.Error
}

func (notebook Notebook) BatchAddNotebook(notebooks *[]Notebook) error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.CreateInBatches(notebooks, len(*notebooks))
	return result.Error
}

func (notebook *Notebook) UpdateNotebook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	dataSet := db.Model(&notebook).Where("NOTE_ID = ?", notebook.NoteId)
	nilTime := time.Time{}
	if notebook.FinishTime != nilTime {
		result := dataSet.Update("FINISH_TIME", notebook.FinishTime)
		if result.Error != nil {
			return result.Error
		}
	}
	if notebook.RealFinishTime != nilTime {
		result := dataSet.Update("REAL_FINISH_TIME", notebook.RealFinishTime)
		if result.Error != nil {
			return result.Error
		}
	}
	if notebook.Content != "" {
		result := dataSet.Update("CONTENT", notebook.Content)
		if result.Error != nil {
			return result.Error
		}
	}
	if notebook.Status != "" {
		result := dataSet.Update("STATUS", notebook.Status)
		if result.Error != nil {
			return result.Error
		}
	}
	if notebook.Level != "" {
		result := dataSet.Update("LEVEL", notebook.Level)
		if result.Error != nil {
			return result.Error
		}
	}
	if notebook.Owner != "" {
		result := dataSet.Update("OWNER", notebook.Owner)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (notebook *Notebook) DeleteNotebook() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&notebook, "NOTE_ID = ?", notebook.NoteId)
	return result.Error
}

func (notebook *Notebook) GetNotebooks(finishTime string, limit int, offset int) ([]Notebook, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	notebooks := make([]Notebook, 0)
	dataSet := db.Model(&notebooks)
	if finishTime != "" && notebook.Status != "" {
		statusArr := strings.Split(notebook.Status, ",")
		dataSet.Where("DATE_FORMAT(FINISH_TIME,'%Y%m%d')<=? AND STATUS IN (?) AND USERNAME = ?", finishTime, statusArr, notebook.Username)
	} else if notebook.Status != "" {
		statusArr := strings.Split(notebook.Status, ",")
		dataSet.Where("STATUS IN (?) AND USERNAME = ?", statusArr, notebook.Username)
	} else if finishTime != "" {
		dataSet.Where("DATE_FORMAT(FINISH_TIME,'%Y%m%d')<=? AND USERNAME = ?", finishTime, notebook.Username)
	} else {
		dataSet.Where("USERNAME = ?", notebook.Username)
	}
	result := dataSet.Limit(limit).Offset(offset).Order("STATUS DESC, FINISH_TIME ASC, LEVEL ASC").Find(&notebooks)
	return notebooks, result.Error
}

func (notebook *Notebook) GetNotebooksCount(finishTime time.Time, status string) (int64, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return 0, errors.New("database connection is nil")
	}
	var count int64
	db = db.Model(&notebook)
	nilTime := time.Time{}
	if finishTime != nilTime {
		db = db.Where("FINISH_TIME=?", finishTime)
	}
	if status != "" {
		db = db.Where("STATUS=?", status)
	}
	result := db.Count(&count)
	log.Debug("count is %d", fmt.Sprint(count))
	return count, result.Error
}
