package goutil

import (
	"github.com/jinzhu/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Sex  int
}

/**
注意是否init()被注释
*/
func TestGetMysqlDB(t *testing.T) {
	db := GetMysqlDB()
	db.CreateTable(&User{})
}

type Site struct {
	gorm.Model
	SiteName string
}

/**
测试OpenMysqlDB请先注释掉init()
*/
func TestOpenMysqlDB(t *testing.T) {
	//userName, password, host, port, DbName, timeOut
	//OpenMysqlDB(100, 10, "root", "123456", "127.0.0.1", "3306", "t_db", "5s")
	dbParameters := []string{
		"root",
		"123456",
		"127.0.0.1",
		"3306",
		"t_db",
		"5s",
	}
	OpenMysqlDB(100, 10, dbParameters...)
	models := []interface{}{
		&User{},
		&Site{},
	}
	db := GetMysqlDB()
	db.CreateTable(models...)
}
