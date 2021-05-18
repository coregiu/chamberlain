package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

/*System configuration*/
var config *ChamberlainConfig

/*Database connection*/
var db *gorm.DB

type ChamberlainConfig struct {
	LogConfig      *LogConfig
	DatabaseConfig *DatabaseConfig
	BlogWorkPath   string
	BlogRepos      *[]*BlogRepository
}

type LogConfig struct {
	Path     string
	LogLevel int8
}

type DatabaseConfig struct {
	Database string
	Host     string
	Port     int
	Username string
	Password string
}

type BlogRepository struct {
	RepoPath string
	RepoName string
}

func init() {
	initLogAndDbConfig()
	initBlogsConfig()
	dbConfig := config.DatabaseConfig
	initDbSource(dbConfig)
}

func initDbSource(dbConfig *DatabaseConfig) {
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

func GetDbConnection() *gorm.DB {
	return db
}

func initLogAndDbConfig() {
	logConfig := new(LogConfig)
	logConfig.Path = "/var/chamberlain.log"
	//Default is Info level = 1
	logConfig.LogLevel = 1

	databaseConfig := new(DatabaseConfig)
	databaseConfig.Database = "chamberlain"
	databaseConfig.Host = "database"
	databaseConfig.Password = "199527"
	databaseConfig.Port = 3306
	databaseConfig.Username = "root"

	config = new(ChamberlainConfig)
	config.LogConfig = logConfig
	config.DatabaseConfig = databaseConfig
}

func GetSystemConfig() *ChamberlainConfig {
	return config
}

func initBlogsConfig() {
	config.BlogWorkPath = "/giu/chamberlain/books"
	repos := make([]*BlogRepository, 2)
	philosophyRepo := new(BlogRepository)
	philosophyRepo.RepoName = "Philosophy"
	philosophyRepo.RepoPath = "https://github.com/coregiu/philosophy.git"

	technologyRepo := new(BlogRepository)
	technologyRepo.RepoName = "Technology"
	technologyRepo.RepoPath = "https://github.com/coregiu/summary.git"

	repos[0] = philosophyRepo
	repos[1] = technologyRepo
	config.BlogRepos = &repos
}

func GetBlogsConfig() (BlogWorkPath string, BlogRepos *[]*BlogRepository) {
	return config.BlogWorkPath, config.BlogRepos
}