package goutil

import (
	"fmt"
	"testing"
)

func TestGetConfigInstance(t *testing.T) {
	c := NewYamlUtil().ReadYamlConfig("./config/dev.yaml")
	fmt.Println(c.JsonCodeDescription)
}
