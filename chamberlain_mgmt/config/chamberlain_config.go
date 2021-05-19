package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*System configuration*/
var globalConfig *ChamberlainConfig

/*Database connection*/
var db *gorm.DB

const DebugLevel = 0
const InfoLevel = 1
const WarnLevel = 2
const ErrorLevel = 3

type (
	ChamberlainConfig struct {
		LogConfig        *LogConfig
		DatabaseConfig   *DatabaseConfig
		BlogWorkPath     string
		BlogRepositories *[]*BlogRepository
	}
	LogConfig struct {
		Path     string
		LogLevel int
	}
	DatabaseConfig struct {
		Database string
		Host     string
		Port     int
		Username string
		Password string
	}
	BlogRepository struct {
		RepoPath string
		RepoName string
	}

	Loader interface {
		loadConfigFromFile() error
		loadDefaultLogConfig()
		loadDefaultDbConfig()
		loadDefaultBlogConfig()
		initDbSource()
	}
)

func init() {
	globalConfig = new(ChamberlainConfig)
	err := globalConfig.loadConfigFromFile()
	if err != nil {
		fmt.Println("failed to load configuration from file")
		globalConfig.loadDefaultLogConfig()
		globalConfig.loadDefaultDbConfig()
		globalConfig.loadDefaultBlogConfig()
	}
	globalConfig.initDbSource()
}

func (chamberlainConfig *ChamberlainConfig) loadConfigFromFile() error {
	workPath, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get work path")
		return err
	}
	fmt.Println(workPath)
	content, ferr := ioutil.ReadFile(workPath + string(os.PathSeparator) + "chamberlain.yml")
	if ferr != nil {
		fmt.Println("failed to read chamberlain yml file")
		return ferr
	}

	contentArr := strings.Split(string(content), "\n")
	fmt.Println(contentArr)

	if -1 == readFileObject(&contentArr, 0, 0, reflect.ValueOf(chamberlainConfig)) {
		return errors.New("failed to read config file")
	}
	return nil
}

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
func readFileObject(contentArr *[]string, startLine int, preSpace int, reflectObject reflect.Value) int {
	fileLines := len(*contentArr)
	if startLine >= fileLines || startLine < 0 {
		fmt.Println("read position out of content array")
		return -1
	}
	iLoop := startLine
	for ; iLoop < fileLines; {
		content := (*contentArr)[iLoop]
		trimContent := strings.Trim(content, " ")
		if strings.Index(trimContent, "#") == 0 || trimContent == "" {
			iLoop++
			continue
		}
		spaceLength := len(content) - len(strings.TrimLeft(content, " "))
		if spaceLength < preSpace {
			fmt.Println("return to parent")
			return iLoop
		} else if spaceLength > preSpace || iLoop == -1 {
			fmt.Println("format error")
			return -1
		} else {
			colonPosition := strings.Index(trimContent, ":")
			key := trimContent[0:colonPosition]
			key = strings.Trim(key, " ")
			fieldValue := reflectObject.Elem().FieldByName(key)

			if colonPosition+1 != len(trimContent) {
				value := trimContent[colonPosition+1:]
				value = strings.Trim(value, " ")
				fmt.Println(key + "=" + value)
				// reflect invoke to set value.
				switch fieldValue.Kind() {
				case reflect.Int:
					valueInt, _ := strconv.ParseInt(value, 10, 64)
					fieldValue.SetInt(valueInt)
					break
				case reflect.String:
					fieldValue.SetString(value)
					break
				}
				iLoop++
			} else {
				if iLoop == fileLines-1 {
					break
				}
				nextContent := (*contentArr)[iLoop+1]
				nextSpaceLength := len(nextContent) - len(strings.TrimLeft(nextContent, " "))
				nextTrimContent := strings.Trim(nextContent, " ")
				// reflect init struct
				if strings.Index(nextTrimContent, "-") != 0 {
					fieldObject := reflect.New(fieldValue.Type().Elem())
					iLoop = readFileObject(contentArr, iLoop+1, nextSpaceLength, fieldObject)
					fieldValue.Set(fieldObject)
				} else {
					fieldArray := reflect.New(fieldValue.Type().Elem())
					iLoop = readFileArray(contentArr, iLoop+1, nextSpaceLength, fieldArray)
					fieldValue.Set(fieldArray)
				}
			}
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
func readFileArray(contentArr *[]string, startLine int, preSpace int, targetArray reflect.Value) int {
	fileLines := len(*contentArr)
	if startLine >= fileLines || startLine < 0 {
		fmt.Println("read position out of content array")
		return -1
	}
	iLoop := startLine
	var fieldObject reflect.Value
	targetElem := targetArray.Elem()
	for ; iLoop < fileLines; {
		content := (*contentArr)[iLoop]
		trimContent := strings.Trim(content, " ")
		if strings.Index(trimContent, "#") == 0 || trimContent == "" {
			iLoop++
			continue
		}

		isArrayStartLine := strings.Index(trimContent, "-") == 0
		spaceLength := len(content) - len(strings.TrimLeft(content, " "))
		if !isArrayStartLine {
			spaceLength -= 2
		}
		if spaceLength < preSpace {
			fmt.Println("return to parent")
			return iLoop
		} else if spaceLength > preSpace || iLoop == -1 {
			fmt.Println("format error")
			return -1
		} else {
			colonPosition := strings.Index(trimContent, ":")
			var key string
			if isArrayStartLine {
				if iLoop != startLine {
					targetElem = reflect.Append(targetElem, fieldObject)
				}
				fieldObject = reflect.New(targetArray.Type().Elem().Elem().Elem())

				key = trimContent[2:colonPosition]
			} else {
				key = trimContent[0:colonPosition]
			}

			key = strings.Trim(key, " ")
			value := trimContent[colonPosition+1:]
			value = strings.Trim(value, " ")
			fieldValue := fieldObject.Elem().FieldByName(key)
			switch fieldValue.Kind() {
			case reflect.Int:
				valueInt, _ := strconv.ParseInt(value, 10, 64)
				fieldValue.SetInt(valueInt)
				break
			case reflect.String:
				fieldValue.SetString(value)
				break
			}
			iLoop++
		}
	}

	if iLoop > startLine {
		targetElem = reflect.Append(targetElem, fieldObject)
	}
	targetArray.Elem().Set(targetElem)
	return iLoop
}

func (chamberlainConfig *ChamberlainConfig) initDbSource() {
	dbConfig := chamberlainConfig.DatabaseConfig
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to create database connection!")
	} else {
		log.Printf("Get db connection successfully!" + fmt.Sprintf("%v", db.Config.AllowGlobalUpdate))
	}
	sqlDb, err := db.DB()
	if err != nil {
		log.Printf("Failed to create database connection!")
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(time.Hour)
}

func (chamberlainConfig *ChamberlainConfig) loadDefaultLogConfig() {
	logConfig := new(LogConfig)
	logConfig.Path = "/var/chamberlain.log"
	//Default is Info level = 1
	logConfig.LogLevel = InfoLevel

	chamberlainConfig.LogConfig = logConfig
}

func (chamberlainConfig *ChamberlainConfig) loadDefaultDbConfig() {
	databaseConfig := new(DatabaseConfig)
	databaseConfig.Database = "chamberlain"
	databaseConfig.Host = "database"
	databaseConfig.Password = "199527"
	databaseConfig.Port = 3306
	databaseConfig.Username = "root"

	chamberlainConfig.DatabaseConfig = databaseConfig
}

func (chamberlainConfig *ChamberlainConfig) loadDefaultBlogConfig() {
	chamberlainConfig.BlogWorkPath = "/giu/chamberlain/books"
	repos := make([]*BlogRepository, 2)
	philosophyRepo := new(BlogRepository)
	philosophyRepo.RepoName = "Philosophy"
	philosophyRepo.RepoPath = "https://github.com/coregiu/philosophy.git"

	technologyRepo := new(BlogRepository)
	technologyRepo.RepoName = "Technology"
	technologyRepo.RepoPath = "https://github.com/coregiu/summary.git"

	repos[0] = philosophyRepo
	repos[1] = technologyRepo
	chamberlainConfig.BlogRepositories = &repos
}

func GetSystemConfig() *ChamberlainConfig {
	return globalConfig
}

func GetBlogsConfig() (BlogWorkPath string, BlogRepos *[]*BlogRepository) {
	return globalConfig.BlogWorkPath, globalConfig.BlogRepositories
}

func GetDbConnection() *gorm.DB {
	return db
}
