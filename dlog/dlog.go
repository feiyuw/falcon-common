package dlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

type Level int

var mutex sync.Mutex

const (
	LevelDebug = Level(iota)
	LevelInfo
	LevelWarning
	LevelError
)

type LogTemplate struct {
	Level    string
	Time     string
	FileName string
	FileLine string
	Tag      string
	Message  string
}

const TIME_FORMAT = "2006-01-02T15:04:05.000Z"
const LOG_FORMAT = "[{{.Time}}][{{.Level}}]<{{.Tag}}> {{.Message}}\n"

type logger struct {
	log_tmpl *template.Template
	lv       Level
	dst      io.Writer
}

type Logger interface {
	Debug(tag string, a ...interface{})
	Debugf(tag string, msg_fmt string, a ...interface{})

	Info(tag string, a ...interface{})
	Infof(tag string, msg_fmt string, a ...interface{})

	Warn(tag string, a ...interface{})
	Warnf(tag string, msg_fmt string, a ...interface{})

	Error(tag string, a ...interface{})
	Errorf(tag string, msg_fmt string, a ...interface{})

	Fatal(tag string, a ...interface{})
	Fatalf(tag string, msg_fmt string, a ...interface{})

	GetInfo() (*template.Template, Level, io.Writer)
}

type debug_logger struct {
	*logger
}

type info_logger struct {
	*debug_logger
}

type warn_logger struct {
	*info_logger
}

type error_logger struct {
	*warn_logger
}

func NewLogger(dst io.Writer, level Level) (l Logger, err error) {
	t, err := template.New("log").Parse(LOG_FORMAT)
	if err != nil {
		return
	}

	l = &logger{
		log_tmpl: t,
		lv:       level,
		dst:      dst,
	}

	if level >= LevelDebug {
		l = &debug_logger{l.(*logger)}
	}

	if level >= LevelInfo {
		l = &info_logger{l.(*debug_logger)}
	}

	if level >= LevelWarning {
		l = &warn_logger{l.(*info_logger)}
	}

	if level >= LevelError {
		l = &error_logger{l.(*warn_logger)}
	}

	return
}

func (l *logger) Debug(tag string, a ...interface{}) {
	l.print(tag, "D", a...)
	return
}

func (l *logger) Debugf(tag string, msg_fmt string, a ...interface{}) {
	l.printf(tag, "D", msg_fmt, a...)
	return
}

func (l *logger) Info(tag string, a ...interface{}) {
	l.print(tag, "I", a...)
	return
}

func (l *logger) Infof(tag string, msg_fmt string, a ...interface{}) {
	l.printf(tag, "I", msg_fmt, a...)
	return
}

func (l *logger) Warn(tag string, a ...interface{}) {
	l.print(tag, "W", a...)
	return
}

func (l *logger) Warnf(tag string, msg_fmt string, a ...interface{}) {
	l.printf(tag, "W", msg_fmt, a...)
	return
}

func (l *logger) Error(tag string, a ...interface{}) {
	l.print(tag, "E", a...)
	return
}

func (l *logger) Errorf(tag string, msg_fmt string, a ...interface{}) {
	l.printf(tag, "E", msg_fmt, a...)
	return
}

func (l *logger) Fatal(tag string, a ...interface{}) {
	l.print(tag, "F", a...)
	os.Exit(1)
}
func (l *logger) Fatalf(tag string, msg_fmt string, a ...interface{}) {
	l.printf(tag, "F", msg_fmt, a...)
	os.Exit(1)
}

func (l *logger) print(tag string, lv string, a ...interface{}) {
	s := ""

	for _, v := range a {
		switch v := v.(type) {
		default:
			s += fmt.Sprintf("%v ", v)
		case error:
			s += fmt.Sprintf("%s ", v.Error())
		}
	}
	l.printer(tag, lv, s)
}

func (l *logger) printf(tag string, lv string, msg_fmt string, a ...interface{}) {
	s := fmt.Sprintf(msg_fmt, a...)
	l.printer(tag, lv, s)
}

func (l *logger) printer(tag string, lv string, str string) {
	_, file_name, line_num, ok := runtime.Caller(3)
	if !ok {
		return
	}

	file_name_s := file_name[strings.LastIndex(file_name, "/")+1:]

	d := &LogTemplate{
		Level:    lv,
		Time:     time.Now().UTC().Format(TIME_FORMAT),
		FileName: file_name_s,
		FileLine: strconv.Itoa(line_num),
		Tag:      tag,
		Message:  str,
	}
	mutex.Lock()
	l.log_tmpl.Execute(l.dst, d)
	mutex.Unlock()
	return
}

func (l *logger) GetInfo() (*template.Template, Level, io.Writer) {
	return l.log_tmpl, l.lv, l.dst
}

func (l *info_logger) Debug(tag string, a ...interface{}) {
	return
}

func (l *info_logger) Debugf(tag string, msg_fmt string, a ...interface{}) {
	return
}

func (l *warn_logger) Info(tag string, a ...interface{}) {
	return
}

func (l *warn_logger) Infof(tag string, msg_fmt string, a ...interface{}) {
	return
}

func (l *error_logger) Warn(tag string, a ...interface{}) {
	return
}

func (l *error_logger) Warnf(tag string, msg_fmt string, a ...interface{}) {
	return
}

func SwitchLevel(old *Logger, level Level) {
	log_tmpl, _, dst := (*old).GetInfo()
	var l Logger
	l = &logger{
		log_tmpl: log_tmpl,
		lv:       level,
		dst:      dst,
	}

	if level >= LevelDebug {
		l = &debug_logger{l.(*logger)}
	}

	if level >= LevelInfo {
		l = &info_logger{l.(*debug_logger)}
	}

	if level >= LevelWarning {
		l = &warn_logger{l.(*info_logger)}
	}

	if level >= LevelError {
		l = &error_logger{l.(*warn_logger)}
	}

	*old = l
}

func StringToLevel(str string) (Level, error) {
	switch strings.ToLower(str) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warning":
		return LevelWarning, nil
	case "error":
		return LevelError, nil
	default:
		return LevelDebug, fmt.Errorf("unknown log level:%s", str)
	}
}
