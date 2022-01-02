package server

import (
	io "chamberlain_mgmt/input"
	"chamberlain_mgmt/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddInputHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		input := &io.Input{}
		inputs := make([]io.Input, 0)
		err := context.BindJSON(&inputs)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("input length =%s", fmt.Sprint(len(inputs)))
		err = input.BatchAddInput(&inputs)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Add input successfully.",
			})
		}
	}
}

func UpdateInputHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		input := io.Input{}
		err := context.BindJSON(&input)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("InputTime =%s", fmt.Sprint(input.InputTime))
		err = input.UpdateInput()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Update input successfully.",
			})
		}
	}
}

func DeleteInputHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		input := io.Input{}
		err := context.BindJSON(&input)
		if err != nil {
			log.Error(err.Error())
			context.String(500, err.Error())
			return
		}
		log.Debug("InputTime =%s", fmt.Sprint(input.InputTime))
		err = input.DeleteInput()
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{
				"result": "Delete input successfully.",
			})
		}
	}
}

func GetInputsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := getIntParam(context, "limit", 10)
		offset := getIntParam(context, "offset", 0)
		year := getUInt16Param(context, "year", 0)
		month := getUInt8Param(context, "month", 0)
		log.Info("limit = %d, offset = %d, year = %d, month = %d", limit, offset, year, month)

		input := new(io.Input)
		inputs, err := input.GetInputs(year, month, limit, offset)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, inputs)
		}
	}
}

func GetInputsCountHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		year := getUInt16Param(context, "year", 0)
		month := getUInt8Param(context, "month", 0)
		log.Info("year = %d, month = %d", year, month)

		input := new(io.Input)
		inputsCount, err := input.GetInputsCount(year, month)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, gin.H{"count": inputsCount})
		}
	}
}

func GetStatisticHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		year := getUInt16Param(context, "year", 0)

		input := new(io.Input)
		inputs, err := input.GetsStatisticsByYear(year)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, inputs)
		}
	}
}

func GetStatisticByMonthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		year := getUInt16Param(context, "year", 0)

		input := new(io.Input)
		inputs, err := input.GetStatisticsByMonth(year)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, inputs)
		}
	}
}

func GetStatisticByTypeHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		year := getUInt16Param(context, "year", 0)

		input := new(io.Input)
		inputs, err := input.GetStatisticsByType(year)
		if err != nil {
			context.String(500, err.Error())
		} else {
			context.JSON(200, inputs)
		}
	}
}

func getIntParam(context *gin.Context, paramName string, defaultValue int) int {
	valueStr, _ := context.GetQuery(paramName)
	if valueStr == "" {
		log.Warn("no parameter %s input", paramName)
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Warn("failed to convert the input parameter %s to int", valueStr)
		return defaultValue
	} else {
		return value
	}
}

func getUInt16Param(context *gin.Context, paramName string, defaultValue uint16) uint16 {
	valueStr, _ := context.GetQuery(paramName)
	if valueStr == "" {
		log.Warn("no parameter %s input", paramName)
		return defaultValue
	}
	value, err := strconv.ParseUint(valueStr, 10, 12)
	if err != nil {
		log.Warn("failed to convert the input parameter %s to int", valueStr)
		return defaultValue
	} else {
		return uint16(value)
	}
}

func getUInt8Param(context *gin.Context, paramName string, defaultValue uint8) uint8 {
	valueStr, _ := context.GetQuery(paramName)
	if valueStr == "" {
		log.Warn("no parameter %s input", paramName)
		return defaultValue
	}
	value, err := strconv.ParseUint(valueStr, 10, 4)
	if err != nil {
		log.Warn("failed to convert the input parameter %s to int", valueStr)
		return defaultValue
	} else {
		return uint8(value)
	}
}