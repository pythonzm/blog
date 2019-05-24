package middleware

import (
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Level int

const (
	INFO Level = iota
	WARNING
	ERROR
	FATAL
)

var (
	file *os.File
	e    error
)

func CustomLogger() gin.HandlerFunc {
	gin.DisableConsoleColor()

	file, e = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if e != nil {
		panic(e)
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	g := gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		levelFlags := []string{"INFO", "WARN", "ERROR", "FATAL"}
		var level string
		status := param.StatusCode

		switch {
		case status > 499:
			level = levelFlags[FATAL]
		case status > 399:
			level = levelFlags[ERROR]
		case status > 299:
			level = levelFlags[WARNING]
		default:
			level = levelFlags[INFO]
		}
		return fmt.Sprintf("[%s] - %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			level,
			param.ClientIP,
			param.TimeStamp.Format(utils.AppInfo.TimeFormat),
			param.Method,
			param.Path,
			param.Request.Proto,
			status,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})

	return g
}

func CloseLogFile() {
	if err := file.Close(); err != nil {
		return
	}
}
