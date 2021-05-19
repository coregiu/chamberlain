package log

import (
	"chamberlain_mgmt/config"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
)

var logLevel int

var (
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	errors  *log.Logger
)

func init() {
	logConfig := config.GetSystemConfig().LogConfig
	file := logConfig.Path
	logLevel = logConfig.LogLevel
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
		return
	}
	defer logFile.Close()
	debug = log.New(io.MultiWriter(os.Stdout, logFile), "Info:", log.Ldate|log.Ltime)
	info = log.New(io.MultiWriter(os.Stderr, logFile), "Info:", log.Ldate|log.Ltime)
	warning = log.New(io.MultiWriter(os.Stdout, logFile), "Warning:", log.Ldate|log.Ltime)
	errors = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime)
}

func Debug(printFormat string, printParams ...interface{}) {
	if logLevel <= config.DebugLevel {
		debug.Printf(printFormat+printCallerName(), printParams)
	}
}

func Info(printFormat string, printParams ...interface{}) {
	if logLevel <= config.InfoLevel {
		info.Printf(printCallerName()+printFormat, printParams)
	}
}

func Warn(printFormat string, printParams ...interface{}) {
	if logLevel <= config.WarnLevel {
		warning.Printf(printCallerName()+printFormat, printParams)
	}
}

func Error(printFormat string, printParams ...interface{}) {
	if logLevel <= config.ErrorLevel {
		errors.Printf(printCallerName()+printFormat, printParams)
	}
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	file := runtime.FuncForPC(pc).Name()
	_, line := runtime.FuncForPC(pc).FileLine(pc)
	return " " + file + ":" + strconv.Itoa(line) + " | "
}