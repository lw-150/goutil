package goutil

import (
	"testing"
)

func TestInitZapLog(t *testing.T) {
	InitZapLog("./logs/spikeProxy1.log", 128, 30, 7, "zap_log")
	type User struct {
		Name string
		Sex  string
		Age  int
	}
	s := User{}
	s.Name = "li wei"
	s.Sex = "man"
	s.Age = 103

	Log.Info(s)
	Logger.Error("3")
	Logger.Sugar().Debug(s)
}
