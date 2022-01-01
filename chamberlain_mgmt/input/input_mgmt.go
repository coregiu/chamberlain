package input

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Input struct {
	InputTime   uint32  `gorm:"column:INPUT_TIME"`
	Year        uint16  `gorm:"column:YEAR"`
	Month       uint8   `gorm:"column:MONTH"`
	Type        string  `gorm:"column:TYPE"`
	Base        float32 `gorm:"column:BASE"`
	AllInput    float32 `gorm:"column:ALL_INPUT"`
	Tax         float32 `gorm:"column:TAX"`
	Actual      float32 `gorm:"column:ACTUAL"`
	Description string  `gorm:"column:DESCRIPTION"`
}

type InputMgmt interface {
	AddInput() error
	UpdateInput() error
	DeleteInput() error
	/*Get details*/
	GetInputs(year uint16, month uint8, limit int, offset int) ([]Input, error)
	/*Get count*/
	GetInputsCount(year uint16, month uint8) (int64, error)
	/*Get sum of month for a year*/
	GetStatisticsByMonth(year uint16) ([]Input, error)
	/*Get sum of type for a year*/
	GetStatisticsByType(year uint16) ([]Input, error)
	/*Get sum of year for all yearthis.inputInfo.s*/
	GetsStatisticsByYear([]Input, error)
}

func (Input) TableName() string  {
	return "INPUTS"
}

func (input *Input) AddInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&input)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to insert input:" + fmt.Sprint(input.InputTime))
	}
	return nil
}

func (input *Input) UpdateInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Model(&Input{}).Where("INPUT_TIME = ?", input.InputTime).
		Update("TYPE", input.Type).
		Update("BASE", input.Base).
		Update("ALL_INPUT", input.AllInput).
		Update("TAX", input.Tax).
		Update("ACTUAL", input.Actual).
		Update("DESCRIPTION", input.Description)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (input *Input) DeleteInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&Input{}, "INPUT_TIME = ?", input.InputTime)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete input:" + fmt.Sprint(input.InputTime))
	}
	return nil
}

func (input *Input) GetInputs(year uint16, month uint8, limit int, offset int) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	if year > 0 {
		db = db.Model(&input).Where("YEAR = ?", year)
	}
	if month > 0 {
		db = db.Model(&input).Where("MONTH = ?", month)
	}
	result := db.Order("INPUT_TIME DESC").Limit(limit).Offset(offset).Find(&inputs)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to get input")
	}
	return inputs, nil
}

func (input *Input) GetInputsCount(year uint16, month uint8) (int64, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return 0, errors.New("database connection is nil")
	}
	var count int64
	db = db.Model(&input)
	if year > 0 {
		db = db.Where("YEAR = ?", year)
	}
	if month > 0 {
		db = db.Where("MONTH = ?", month)
	}
	result := db.Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed to get input")
	}
	log.Debug("count is %d", fmt.Sprint(count))
	return count, nil
}

func (input *Input) GetStatisticsByMonth(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	result = db.Raw("SELECT YEAR, MONTH, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ? GROUP BY YEAR, MONTH", year).Scan(&inputs)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		log.Warn("no result go get.")
		return nil, errors.New("failed to get input")
	}
	return inputs, nil
}

func (input *Input) GetStatisticsByType(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	result = db.Raw("SELECT YEAR, TYPE, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ? GROUP BY YEAR, TYPE", year).Scan(&inputs)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to get input")
	}
	return inputs, nil
}

func (input *Input) GetsStatisticsByYear(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	if year > 0 {
		result = db.Raw("SELECT YEAR, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS GROUP BY YEAR").Scan(&inputs)
	} else {
		result = db.Raw("SELECT YEAR, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ?", year).Scan(&inputs)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to get input")
	}
	return inputs, nil
}
