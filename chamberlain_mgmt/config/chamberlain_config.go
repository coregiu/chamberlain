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
		LogLevel int `DEBUG:"0" INFO:"1" WARN:"2" ERROR:"3"`
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
	fmt.Printf("work path=%s\n", workPath)
	content, err := ioutil.ReadFile(workPath + string(os.PathSeparator) + "chamberlain.yml")
	if err != nil {
		fmt.Println("failed to read chamberlain yml file")
		return err
	}

	contentArr := strings.Split(string(content), "\n")

	if -1 == ReadFileObject(&contentArr, 0, 0, reflect.ValueOf(chamberlainConfig)) {
		return errors.New("failed to read config file")
	}
	return nil
}

func (chamberlainConfig *ChamberlainConfig) initDbSource() {
	dbConfig := chamberlainConfig.DatabaseConfig
	connectArr := []string{dbConfig.Username, ":", dbConfig.Password, "@tcp(", dbConfig.Host, ":", strconv.Itoa(dbConfig.Port), ")/", dbConfig.Database, "?charset=utf8mb4&parseTime=True&loc=Local"}
	dbUrl := strings.Join(connectArr, "")
	var err error
	db, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to create database connection!")
		return
	}

	log.Printf("Get db connection successfully!")
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
	databaseConfig.Password = ""
	databaseConfig.Port = 3306
	databaseConfig.Username = ""

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
