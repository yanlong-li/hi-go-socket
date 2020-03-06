package logger

import (
	"fmt"
)

//日志结构体
type Log struct {
	code  uint
	msg   string
	level uint8
	data  []interface{}
}

const (
	ALL uint8 = iota
	TRACE
	DEBUG
	INFO
	WARNING
	FATAL
)

func create(level uint8, msg string, code uint, data ...interface{}) *Log {
	log := new(Log)
	log.level = level
	log.msg = msg
	log.code = code
	log.data = data
	return log
}

//所有日志
func All(msg string, code uint, data ...interface{}) {
	log := create(DEBUG, msg, code, data...)
	log.handel()
}

//跟踪日志
func Trace(msg string, code uint, data ...interface{}) {
	log := create(TRACE, msg, code, data...)
	log.handel()
}

//调试日志
func Debug(msg string, code uint, data ...interface{}) {
	log := create(DEBUG, msg, code, data...)
	log.handel()
}

//信息日志
func Info(msg string, code uint, data ...interface{}) {
	log := create(INFO, msg, code, data...)
	log.handel()
}

//警告日志
func Warning(msg string, code uint, data ...interface{}) {
	log := create(WARNING, msg, code, data...)
	log.handel()
}

//致命日志
func Fatal(msg string, code uint, data ...interface{}) {
	log := create(FATAL, msg, code, data...)
	log.handel()
}

func (log *Log) handel() {

	fmt.Printf("[%s][%d] %s \n", GetLabel(log.level), log.code, log.msg)
	if len(log.data) > 0 {
		fmt.Println(log.data)
	}

}

func GetLabel(levelType uint8) string {

	level := ""
	switch levelType {
	case ALL:
		level = "ALL"
	case TRACE:
		level = "TRACE"
	case DEBUG:
		level = "DEBUG"
	case INFO:
		level = "INFO"
	case WARNING:
		level = "WARNING"
	case FATAL:
		level = "FATAL"
	}
	return level
}
