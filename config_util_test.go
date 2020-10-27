package goutil

import (
	"fmt"
	"testing"
)

func TestGetConfigInstance(t *testing.T) {
	c := GetConfigInstance("./config/dev.yaml")
	fmt.Println(c.Mysql.Password)
}
