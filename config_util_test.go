package goutil

import (
	"fmt"
	"testing"
)

func TestGetConfigInstance(t *testing.T) {
	GetConfigInstance("./config/dev.yaml")
	fmt.Println(Config.Mysql.Password)
}
