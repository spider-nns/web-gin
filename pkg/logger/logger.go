package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}
func (l *Logger) WithFields(f Fields) *Logger {
	cl := l.clone()
	if cl.fields == nil {
		cl.fields = make(Fields)
	}
	for k, v := range f {
		cl.fields[k] = v
	}
	return cl
}
func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	cl := l.clone()
	cl.ctx = ctx
	return cl
}

func (l *Logger) WithCaller(skip int) *Logger {
	cl := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		cl.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return cl
}
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	cl := l.clone()
	cl.callers = callers
	return cl
}

func (l *Logger) JsonFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JsonFormat(level, message))
	content := string(body)
	switch level {
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)

	}
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(LevelInfo, fmt.Sprint(v...))
}
func (l *Logger) InfoF(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WithLevel(LevelWarn).Output(LevelWarn, fmt.Sprint(v...))
}
func (l *Logger) WarnF(format string, v ...interface{}) {
	l.WithLevel(LevelWarn).Output(LevelWarn, fmt.Sprintf(format, v...))
}
func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(LevelFatal, fmt.Sprint(v...))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(LevelFatal, fmt.Sprintf(format, v...))
}
func (l *Logger) Error(v ...interface{}) {
	l.WithLevel(LevelError).Output(LevelError, fmt.Sprint(v...))
}
func (l *Logger) ErrorF(format string, v ...interface{}) {
	l.WithLevel(LevelError).Output(LevelError, fmt.Sprintf(format, v...))
}
func (l *Logger) Panic(v ...interface{}) {
	l.WithLevel(LevelPanic).Output(LevelPanic, fmt.Sprint(v...))
}
func (l *Logger) PanicF(format string, v ...interface{}) {
	l.WithLevel(LevelPanic).Output(LevelPanic, fmt.Sprintf(format, v...))
}
