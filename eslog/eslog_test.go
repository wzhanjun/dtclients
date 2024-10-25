package eslog

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Std().Debug("this is debug")
	Std().Info("this is info")
	Std().Warn("this is warn")
	Std().Error("this is error")
	time.Sleep(time.Second * 3)
}
