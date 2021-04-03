package log

import (
	"chamberlain_mgmt/config"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
)

const DEBUG_LEVEL = 0
const INFO_LEVEL = 1
const WARN_LEVEL = 2
const ERROR_LEVEL = 3

var logLevel int8

var (
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
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
	debug = log.New(io.MultiWriter(logFile), "Info:", log.Ldate|log.Ltime)
	info = log.New(io.MultiWriter(logFile), "Info:", log.Ldate|log.Ltime)
	warning = log.New(io.MultiWriter(os.Stdout, logFile), "Warning:", log.Ldate|log.Ltime)
	error = log.New(io.MultiWriter(os.Stderr, logFile), "Error:", log.Ldate|log.Ltime)
}

func Debug(printFormat string, printParams ...string) {
	if logLevel <= DEBUG_LEVEL {
		debug.Printf(printFormat+printCallerName(), printParams)
	}
}

func Info(printFormat string, printParams ...string) {
	if logLevel <= INFO_LEVEL {
		info.Printf(printCallerName()+printFormat, printParams)
	}
}

func Warn(printFormat string, printParams ...string) {
	if logLevel <= WARN_LEVEL {
		warning.Printf(printCallerName()+printFormat, printParams)
	}
}

func Error(printFormat string, printParams ...string) {
	if logLevel <= ERROR_LEVEL {
		error.Printf(printCallerName()+printFormat, printParams)
	}
}

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	file := runtime.FuncForPC(pc).Name()
	_, line := runtime.FuncForPC(pc).FileLine(pc)
	return " " + file + ":" + strconv.Itoa(line) + " | "
}
