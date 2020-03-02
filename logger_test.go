package gologger_test

import (
	"testing"

	"github.com/savsgio/go-logger"
	"go.uber.org/zap"
)

// TODO: Incomplete tests. No real stdout/stderr asserts

func init() {
	log, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(log)
}

func newLogger(lvl string) *gologger.Logger {
	return gologger.New("test", lvl, nil)
}

func TestLogger_Printf(t *testing.T) {
	log := newLogger(gologger.DEBUG)
	log.Printf("golog PRINT message %q", "test ok")
}

func TestLogger_Debug(t *testing.T) {
	log := newLogger(gologger.DEBUG)
	log.Debugf("golog DEBUG message %q", "test ok")
}

func TestLogger_Info(t *testing.T) {
	log := newLogger(gologger.DEBUG)
	log.Infof("golog INFO message %q", "test ok")
}

func TestLogger_Warning(t *testing.T) {
	log := newLogger(gologger.DEBUG)
	log.Warningf("golog WARNING message %q", "test ok")
}

func TestLogger_Errorf(t *testing.T) {
	log := newLogger(gologger.DEBUG)
	log.Errorf("golog ERROR message %q", "test ok")
}
