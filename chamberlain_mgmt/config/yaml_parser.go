package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type YmlDataType int

const (
	BaseType YmlDataType = iota
	Struct
	Array
	ArrayStruct
	ArrayMapFst
	ArrayMap
	ArrayString
	NilValue
	NoData
	Parent
	FormatError
)

/**
 * read config as object.
 *
 * contentArr - file content array by line
 * startLine - parse line start position
 * preSpace - the space of line prefix
 * targetObject - convert target, output parameter
 *
 * return the end line of this object.
 */
func ReadFileObject(contentArr *[]string, startLine int, preSpace int, reflectObject reflect.Value) int {
	fileLines := len(*contentArr)
	iLoop := startLine
	for ; iLoop < fileLines; {
		dataType := getCurrentDataType(contentArr, iLoop, preSpace, false)

		switch dataType {
		case FormatError:
			fmt.Println("format error")
			return -1
		case Parent:
			fmt.Println("return to parent")
			return iLoop
		case NoData:
			iLoop++
			break
		case BaseType:
			key, value := getBaseKeyValue(contentArr, iLoop)
			fieldValue := reflectObject.Elem().FieldByName(key)
			setReflectValue(fieldValue, value)
			iLoop++
			break
		case Struct:
			key := getReflectKey(contentArr, iLoop)
			fieldValue := reflectObject.Elem().FieldByName(key)
			fieldObject := reflect.New(fieldValue.Type().Elem())
			nextSpaceLength := getLineSpaceLength(contentArr, iLoop+1, false)
			iLoop = ReadFileObject(contentArr, iLoop+1, nextSpaceLength, fieldObject)
			fieldValue.Set(fieldObject)
			break
		case Array:
			key := getReflectKey(contentArr, iLoop)
			fieldValue := reflectObject.Elem().FieldByName(key)
			fieldArray := reflect.New(fieldValue.Type().Elem())
			nextSpaceLength := getLineSpaceLength(contentArr, iLoop+1, false)
			iLoop = ReadFileArray(contentArr, iLoop+1, nextSpaceLength, fieldArray)
			fieldValue.Set(fieldArray)
			break
		case NilValue:
			break
		}
	}
	return iLoop
}

/**
 * read config as array.
 *
 * contentArr - file content array by line
 * startLine - parse line start position
 * preSpace - the space of line prefix
 * targetArray - convert target, output parameter
 *
 * return the end line of this object.
 */
func ReadFileArray(contentArr *[]string, startLine int, preSpace int, targetArray reflect.Value) int {
	fileLines := len(*contentArr)
	iLoop := startLine
	var fieldObject reflect.Value
	targetElem := targetArray.Elem()
	for ; iLoop < fileLines; {
		dataType := getCurrentDataType(contentArr, iLoop, preSpace, true)

		switch dataType {
		case FormatError:
			fmt.Println("format error")
			return -1
		case Parent:
			fmt.Println("return to parent")
			return iLoop
		case NoData:
			iLoop++
			break
		case ArrayString:
			iLoop++
			break
		case ArrayStruct:
			iLoop++
			break
		case ArrayMapFst:
			content := (*contentArr)[iLoop]
			trimContent := strings.Trim(content, " ")
			colonPosition := strings.Index(trimContent, ":")
			if iLoop != startLine {
				targetElem = reflect.Append(targetElem, fieldObject)
			}
			fieldObject = reflect.New(targetArray.Type().Elem().Elem().Elem())

			key := trimContent[2:colonPosition]
			key = strings.Trim(key, " ")
			value := trimContent[colonPosition+1:]
			value = strings.Trim(value, " ")
			fieldValue := fieldObject.Elem().FieldByName(key)
			setReflectValue(fieldValue, value)
			iLoop++
			break
		case ArrayMap:
			content := (*contentArr)[iLoop]
			trimContent := strings.Trim(content, " ")
			colonPosition := strings.Index(trimContent, ":")
			key := trimContent[0:colonPosition]
			key = strings.Trim(key, " ")
			value := trimContent[colonPosition+1:]
			value = strings.Trim(value, " ")
			fieldValue := fieldObject.Elem().FieldByName(key)
			setReflectValue(fieldValue, value)
			iLoop++
			break
		case NilValue:
			break
		}
	}

	if iLoop > startLine {
		targetElem = reflect.Append(targetElem, fieldObject)
	}
	targetArray.Elem().Set(targetElem)
	return iLoop
}

func getBaseKeyValue(contentArr *[]string, iLoop int) (string, string) {
	content := (*contentArr)[iLoop]
	trimContent := strings.Trim(content, " ")
	colonPosition := strings.Index(trimContent, ":")
	key := trimContent[0:colonPosition]
	key = strings.Trim(key, " ")
	value := trimContent[colonPosition+1:]
	value = strings.Trim(value, " ")
	return key, value
}

func getLineSpaceLength(contentArr *[]string, iLoop int, isArrayDataLine bool) int {
	content := (*contentArr)[iLoop]
	spaceLength := len(content) - len(strings.TrimLeft(content, " "))
	if isArrayDataLine {
		spaceLength -= 2
	}
	return spaceLength
}

func getReflectKey(contentArr *[]string, iLoop int) string {
	content := (*contentArr)[iLoop]
	trimContent := strings.Trim(content, " ")
	colonPosition := strings.Index(trimContent, ":")
	key := trimContent[0:colonPosition]
	key = strings.Trim(key, " ")
	return key
}

func getCurrentDataType(contentArr *[]string, startLine int, preSpace int, isInArray bool) YmlDataType {
	fileLines := len(*contentArr)
	if startLine >= fileLines || startLine < 0 {
		fmt.Println("read position out of content array")
		return FormatError
	}
	currentContent := (*contentArr)[startLine]
	trimContent := strings.Trim(currentContent, " ")
	if strings.Index(trimContent, "#") == 0 || trimContent == "" {
		startLine++
		return NoData
	}
	isArrayStartLine := strings.Index(trimContent, "-") == 0
	spaceLength := getLineSpaceLength(contentArr, startLine, !isArrayStartLine && isInArray)
	if spaceLength < preSpace {
		fmt.Println("return to parent")
		return Parent
	} else if spaceLength > preSpace || startLine == -1 {
		fmt.Println("format error")
		return FormatError
	} else {
		colonPosition := strings.Index(trimContent, ":")
		if !isInArray && colonPosition+1 != len(trimContent) {
			return BaseType
		} else if !isInArray && colonPosition+1 == len(trimContent) {
			if startLine == fileLines-1 {
				return NilValue
			}
			nextContent := (*contentArr)[startLine+1]
			nextTrimContent := strings.Trim(nextContent, " ")
			if strings.Index(nextTrimContent, "-") == -1 {
				return Struct
			} else {
				return Array
			}
		} else if isInArray && isArrayStartLine {
			return ArrayMapFst
		} else if isInArray && !isArrayStartLine {
			return ArrayMap
		} else {
			return FormatError
		}
	}
}

func setReflectValue(fieldValue reflect.Value, value string) {
	switch fieldValue.Kind() {
	case reflect.Int:
		valueInt, _ := strconv.ParseInt(value, 10, 64)
		fieldValue.SetInt(valueInt)
		break
	case reflect.String:
		if strings.Index(value, "${") != 0 {
			fieldValue.SetString(value)
		} else {
			passVar := value[2 : len(value)-1]
			fieldValue.SetString(os.Getenv(passVar))
		}
		break
	}
}
