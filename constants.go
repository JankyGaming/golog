package golog

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

//LogClient is the log client you will use to write logs to your db or stdout
type LogClient interface {
	//Log t is the type of log, l is the log level, desc is a description of the error, and data is an optional map of data
	Log(t logType, l logLevel, desc string, data map[string]interface{})
}

//formatted strings
const (
	logTypeLogStr   = "| Type: Log   |"
	logTypeErrorStr = "| Type: Error |"

	logLevelExpectedStr = " Level: Expected |"
	logLevelSevereStr   = " Level: Severe   |"
)

type logType string
type logLevel string

const (
	//Log Log type for when you went to simply log something to your client
	Log logType = "Log"
	//Error Log type for when you went to log an error to your client
	Error logType = "Error"

	//Expected Log level for normal logs, the typical log level
	Expected logLevel = "Expected"
	//Severe Log level for errors that need to be evaluated
	Severe logLevel = "Severe"
)

type log struct {
	Description string                 `json:"description" bson:"description"`
	LogType     logType                `json:"logType" bson:"logType"`
	LogLevel    logLevel               `json:"logLevel" bson:"logLevel"`
	Location    string                 `json:"location" bson:"location"`
	Timestamp   time.Time              `json:"timestamp" bson:"timestamp"`
	Data        map[string]interface{} `json:"data" bson:"data"`
	StdOutPrint string                 `json:"-" bson:"-"`
}

func buildLog(t logType, l logLevel, desc string, data map[string]interface{}) log {
	curTime := time.Now()
	pc, _, line, _ := runtime.Caller(2)
	locStr := runtime.FuncForPC(pc).Name() + " line " + strconv.Itoa(line)

	var typeStr string
	var levelStr string

	switch t {
	case Log:
		typeStr = logTypeLogStr
	case Error:
		typeStr = logTypeErrorStr
	}

	switch l {
	case Expected:
		levelStr = logLevelExpectedStr
	case Severe:
		levelStr = logLevelSevereStr
	}

	newLog := log{
		Description: desc,
		LogType:     t,
		LogLevel:    l,
		Location:    locStr,
		Timestamp:   curTime,
		Data:        data,
		StdOutPrint: fmt.Sprintln(curTime.Format("2006-01-02 15:04:05"), typeStr+levelStr, locStr, "|", desc),
	}

	return newLog
}
