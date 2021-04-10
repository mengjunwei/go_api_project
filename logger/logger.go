package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

const logChanLen = 4096

func SetLog() error {
	logsLevel := logs.LevelError
	logLevelStr := beego.AppConfig.DefaultString("LogLevel", "INFO")
	logPath := beego.AppConfig.DefaultString("LogPath", "./logs/")
	logAdapter := beego.AppConfig.DefaultString("LogAdapter", logs.AdapterConsole)
	logMaxsize := beego.AppConfig.DefaultInt("LogMaxsize", 512)
	logMaxDay := beego.AppConfig.DefaultInt("LogMaxDay", 7)

	{
		l := strings.ToLower(logLevelStr)
		switch l {
		case "debug":
			logsLevel = logs.LevelDebug
		case "info":
			logsLevel = logs.LevelInfo
		case "warn":
			logsLevel = logs.LevelWarn
		case "error":
			logsLevel = logs.LevelError
		default:
		}
	}
	logConf := make(map[string]interface{})
	logConf["level"] = logsLevel
	logConf["maxsize"] = logMaxsize
	logConf["maxdays"] = logMaxDay
	logConf["daily"] = true
	logConf["color"] = true

	fn := func(map[string]interface{}) (string, error) {
		confBytes, err := json.Marshal(logConf)
		if err != nil {
			fmt.Println("conf marshal failed,err:", err)
			return "", err
		}
		return string(confBytes), nil
	}

	switch logAdapter {
	case logs.AdapterConsole:
		confStr, err := fn(logConf)
		if err != nil {
			return err
		}
		err = logs.SetLogger(logs.AdapterConsole, confStr)
		if err != nil {
			return err
		}
	case logs.AdapterFile:
		filePath := path.Join(logPath, "log.log")
		logConf["filename"] = filePath
		confStr, err := fn(logConf)
		if err != nil {
			return err
		}
		err = logs.SetLogger(logs.AdapterFile, confStr)
		if err != nil {
			return err
		}
	default:
		return errors.New("conf err:AdapterConsole AdapterFile")
	}
	logs.EnableFullFilePath(true)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(2)
	logs.Async(logChanLen)
	return nil
}

// Critical logs a message at critical level.
func Critical(f interface{}, v ...interface{}) {
	logs.Critical(f, v)
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v)
}

// Warning logs a message at warning level.
func Warning(f interface{}, v ...interface{}) {
	logs.Warning(f, v)
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v)
}

// Info compatibility alias for Warning()
func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v)
}
