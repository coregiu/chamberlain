package log

import (
	"chamberlain_mgmt/config"
	"errors"
	"fmt"
	"time"
)

type SysLog struct {
	LogId       int64     `gorm:"column:LOG_ID"`
	Username    string    `gorm:"column:USERNAME"`
	Operation   string    `gorm:"column:OPERATION"`
	OpTime      time.Time `gorm:"column:OP_TIME"`
	OpResult    string    `gorm:"column:OP_RESULT"`
	Description string    `gorm:"column:DESCRIPTION"`
}

type SysLogMgmt interface {
	RecordOperation(username string, operation string, opResult string, description string) error
	AddSyslog() error
	BatchAddSyslog(inputs *[]SysLog) error
	UpdateSyslog() error
	DeleteSyslog(syslogs *[]SysLog) error
	/*Get details*/
	GetSyslog(username string, operation string, limit int, offset int) ([]SysLog, error)
	/*Get count*/
	GetSyslogCount(username string, operation string) (int64, error)
}

func (SysLog) TableName() string {
	return "LOGS"
}

func (syslog *SysLog) AddSyslog() error {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&syslog)
	return result.Error
}

func (syslog *SysLog) RecordOperation(username string, operation string, opResult string, description string) error {
	syslog.OpTime = time.Now()
	syslog.LogId = syslog.OpTime.Unix()
	syslog.Operation = operation
	syslog.OpResult = opResult
	syslog.Description = description
	syslog.Username = username
	return syslog.AddSyslog()
}

func (syslog SysLog) BatchAddSyslog(syslogs *[]SysLog) error {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.CreateInBatches(syslogs, len(*syslogs))
	return result.Error
}

func (syslog *SysLog) UpdateSyslog() error {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Model(&syslog).Where("LOG_ID = ?", syslog.LogId).
		Update("USERNAME", syslog.Username).
		Update("OPERATION", syslog.Operation).
		Update("OP_RESULT", syslog.OpResult).
		Update("DESCRIPTION", syslog.Description)
	return result.Error
}

func (syslog *SysLog) DeleteSyslog(syslogs *[]SysLog) error {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	if syslogs == nil || len(*syslogs) == 0{
		Warn("no logs to delete")
		return nil
	}
	var selectIds []int64
	for _, log := range *syslogs {
		selectIds = append(selectIds, log.LogId)
	}

	result := db.Delete(&syslog, "LOG_ID IN (?)", selectIds)
	return result.Error
}

func (syslog *SysLog) GetSyslog(username string, operation string, limit int, offset int) ([]SysLog, error) {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	syslogs := make([]SysLog, 0)
	dataSet := db.Model(&syslogs)
	if username != "" {
		dataSet.Where("USERNAME=?", username)
	}
	if operation != "" {
		dataSet.Where("OPERATION=?", operation)
	}
	result := dataSet.Limit(limit).Offset(offset).Order("OP_TIME DESC").Find(&syslogs)
	return syslogs, result.Error
}

func (syslog *SysLog) GetSyslogCount(username string, operation string) (int64, error) {
	db := config.GetDbConnection()
	if db == nil {
		Error("Db connection is nil")
		return 0, errors.New("database connection is nil")
	}
	var count int64
	dataSet := db.Model(&count)
	if username != "" {
		dataSet.Where("USERNAME=?", username)
	}
	if operation != "" {
		dataSet.Where("OPERATION=?", operation)
	}
	result := dataSet.Count(&count)
	Debug("count is %d", fmt.Sprint(count))
	return count, result.Error
}
