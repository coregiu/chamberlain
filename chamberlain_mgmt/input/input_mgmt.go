package input

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Input struct {
	InputTime   uint32
	Year        uint16
	Month       uint8
	Type        string
	Base        float32
	AllInput    float32
	Tax         float32
	Actual      float32
	Description string
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
	/*Get sum of year for all years*/
	GetsStatisticsByYear([]Input, error)
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
	result := db.Model(&Input{}).Where("InputTime = ?", input.InputTime).
		Update("Year", input.Year).
		Update("Month", input.Month).
		Update("Type", input.Type).
		Update("Month", input.Month).
		Update("Base", input.Base).
		Update("AllInput", input.AllInput).
		Update("Tax", input.Tax).
		Update("Actual", input.Actual).
		Update("Description", input.Description)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update input:" + fmt.Sprint(input.InputTime))
	}
	return nil
}

func (input *Input) DeleteInput() error {
	db := config.GetDbConnection()
	if db == nil {
		log.Error("Db connection is nil")
		return errors.New("database connection is nil")
	}
	result := db.Delete(&Input{}, "InputTime = ?", input.InputTime)
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
		db = db.Model(&input).Where("Year = ?", year)
	}
	if month > 0 {
		db = db.Model(&input).Where("Month = ?", month)
	}
	result := db.Limit(limit).Offset(offset).Find(&inputs)
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
		db = db.Where("Year = ?", year)
	}
	if month > 0 {
		db = db.Where("Month = ?", month)
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
	result = db.Raw("select Year, Month, sum(AllInput) AllInput, sum(Tax) Tax, sum(Actual) Actual from inputs where year = ? group by Year, Month", year).Scan(&inputs)
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
	result = db.Raw("select Year, Type, sum(AllInput) AllInput, sum(Tax) Tax, sum(Actual) Actual from inputs where year = ? group by Year, Type", year).Scan(&inputs)

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
		result = db.Raw("select Year, sum(AllInput) AllInput, sum(Tax) Tax, sum(Actual) Actual from inputs group by Year").Scan(&inputs)
	} else {
		result = db.Raw("select Year, sum(AllInput) AllInput, sum(Tax) Tax, sum(Actual) Actual from inputs where Year = ?", year).Scan(&inputs)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("failed to get input")
	}
	return inputs, nil
}
