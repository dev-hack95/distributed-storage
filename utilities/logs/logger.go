package logs

import (
	log "github.com/jeanphorn/log4go"
)

func Error(arg0 interface{}, args ...interface{}) {
	log.Error(arg0, args...)
}

func Info(arg0 interface{}, args ...interface{}) {
	log.Info(arg0, args...)
}

func Debug(arg0 interface{}, args ...interface{}) {
	log.Debug(arg0, args...)
}
