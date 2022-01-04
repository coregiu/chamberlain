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
	BatchAddInput(inputs *[]Input) error
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

func (Input) TableName() string {
	return "INPUTS"
}

func (input *Input) AddInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Create(&input)
	return result.Error
}

func (input Input) BatchAddInput(inputs *[]Input) error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.CreateInBatches(inputs, len(*inputs))
	return result.Error
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
	return result.Error
}

func (input *Input) DeleteInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&Input{}, "INPUT_TIME = ?", input.InputTime)
	return result.Error
}

func (input *Input) GetInputs(year uint16, month uint8, limit int, offset int) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	if year > 0 {
		result := db.Raw("SELECT INPUT_TIME, YEAR, MONTH, TYPE, ROUND(ALL_INPUT, 2) ALL_INPUT, ROUND(TAX, 2) TAX, ROUND(ACTUAL, 2) ACTUAL, BASE, DESCRIPTION FROM INPUTS WHERE YEAR = ? ORDER BY INPUT_TIME DESC", year).Scan(&inputs)
		return inputs, result.Error
	}
	result := db.Raw("SELECT INPUT_TIME, YEAR, MONTH, TYPE, ROUND(ALL_INPUT, 2) ALL_INPUT, ROUND(TAX, 2) TAX, ROUND(ACTUAL, 2) ACTUAL, BASE, DESCRIPTION FROM INPUTS ORDER BY INPUT_TIME DESC", year).Scan(&inputs)
	return inputs, result.Error
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
	log.Debug("count is %d", fmt.Sprint(count))
	return count, result.Error
}

func (input *Input) GetStatisticsByMonth(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	result = db.Raw("SELECT MONTH, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ? GROUP BY MONTH ORDER BY MONTH", year).Scan(&inputs)
	return inputs, result.Error
}

func (input *Input) GetStatisticsByType(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	result = db.Raw("SELECT TYPE, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ? GROUP BY TYPE ORDER BY TYPE", year).Scan(&inputs)

	return inputs, result.Error
}

func (input *Input) GetsStatisticsByYear(year uint16) ([]Input, error) {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return nil, errors.New("database connection is nil")
	}
	inputs := make([]Input, 0)
	var result *gorm.DB
	if year == 0 {
		result = db.Raw("SELECT YEAR, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS GROUP BY YEAR ORDER BY YEAR").Scan(&inputs)
	} else {
		result = db.Raw("SELECT YEAR, SUM(ALL_INPUT) ALL_INPUT, SUM(TAX) TAX, SUM(ACTUAL) ACTUAL FROM INPUTS WHERE YEAR = ? ORDER BY YEAR", year).Scan(&inputs)
	}

	return inputs, result.Error
}
