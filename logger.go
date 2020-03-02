package gologger

import (
	"io"

	"go.uber.org/zap"
)

type Logger struct {
	zap   *zap.Logger
	log   *zap.SugaredLogger
	debug bool
}

// New create new instance of Logger
func New(name string, level string, _ io.Writer) *Logger {
	l := zap.L().
		WithOptions(zap.AddCallerSkip(1)).
		With(zap.String("golog", name))
	return &Logger{
		zap:   l,
		log:   l.Sugar(),
		debug: level == DEBUG,
	}
}

func (l *Logger) SetOutput(io.Writer) { /* nope */ }

func (l *Logger) DebugEnabled() bool {
	return l.debug
}

func (l *Logger) Fatal(msg ...interface{}) {
	l.log.Fatal(msg...)
}

func (l *Logger) Fatalf(msg string, v ...interface{}) {
	l.log.Fatalf(msg, v...)
}

func (l *Logger) Error(msg ...interface{}) {
	l.log.Error(msg...)
}

func (l *Logger) Errorf(msg string, v ...interface{}) {
	l.log.Errorf(msg, v...)
}

func (l *Logger) Warning(msg ...interface{}) {
	l.log.Warn(msg...)
}

func (l *Logger) Warningf(msg string, v ...interface{}) {
	l.log.Warnf(msg, v...)
}

func (l *Logger) Info(msg ...interface{}) {
	l.log.Info(msg...)
}

func (l *Logger) Infof(msg string, v ...interface{}) {
	l.log.Infof(msg, v...)
}

func (l *Logger) Debug(msg ...interface{}) {
	l.log.Debug(msg...)
}

func (l *Logger) Debugf(msg string, v ...interface{}) {
	l.log.Debugf(msg, v...)
}

func (l *Logger) Print(msg ...interface{}) {
	l.log.Debug(msg)
}

func (l *Logger) Printf(msg string, v ...interface{}) {
	l.log.Debugf(msg, v...)
}
