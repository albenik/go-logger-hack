package gologger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	FATAL   = "fatal"
	ERROR   = "error"
	WARNING = "warning"
	INFO    = "info"
	DEBUG   = "debug"
)

var levelMap = map[string]zapcore.Level{
	FATAL:   zapcore.FatalLevel,
	ERROR:   zapcore.ErrorLevel,
	WARNING: zapcore.WarnLevel,
	INFO:    zapcore.InfoLevel,
	DEBUG:   zapcore.DebugLevel,
}

func stringToLevel(l string) zap.AtomicLevel {
	if zl, ok := levelMap[strings.ToLower(l)]; ok {
		return zap.NewAtomicLevelAt(zl)
	}
	panic(fmt.Errorf("unknown log level %q", l))
}
