package logger

import (
	"github.com/feiyuw/falcon-common/dlog"
	"log"
	"os"
)

var (
	defaultLogger dlog.Logger
	logLevel      = "debug"
)

func init() {
	l, err := dlog.StringToLevel(logLevel)
	defaultLogger, err = dlog.NewLogger(os.Stdout, l)
	if err != nil {
		log.Fatalln("init logger error")
	}
}

func SetLogLevel(level string) error {
	l, err := dlog.StringToLevel(level)
	if err != nil {
		return err
	}
	dlog.SwitchLevel(&defaultLogger, l)
	logLevel = level

	return nil
}
func GetLogLevel() string {
	return logLevel
}

func Debug(tag string, a ...interface{}) {
	defaultLogger.Debug(tag, a...)
}

func Debugf(tag string, msg_fmt string, a ...interface{}) {
	defaultLogger.Debugf(tag, msg_fmt, a...)
}

func Info(tag string, a ...interface{}) {
	defaultLogger.Info(tag, a...)
}

func Infof(tag string, msg_fmt string, a ...interface{}) {
	defaultLogger.Infof(tag, msg_fmt, a...)
}

func Warn(tag string, a ...interface{}) {
	defaultLogger.Warn(tag, a...)
}

func Warnf(tag string, msg_fmt string, a ...interface{}) {
	defaultLogger.Warnf(tag, msg_fmt, a...)
}

func Error(tag string, a ...interface{}) {
	defaultLogger.Error(tag, a...)
}
func Errorf(tag string, msg_fmt string, a ...interface{}) {
	defaultLogger.Errorf(tag, msg_fmt, a...)
}

func Fatal(tag string, a ...interface{}) {
	defaultLogger.Fatal(tag, a...)
}
func Fatalf(tag string, msg_fmt string, a ...interface{}) {
	defaultLogger.Fatalf(tag, msg_fmt, a...)
}
