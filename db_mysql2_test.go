package goutil

import (
	"gorm.io/gorm"
	"testing"
)

type User2 struct {
	gorm.Model
	Name string
	Sex  int
}

/**
注意是否init()被注释
*/
func TestGetMysqlDB2(t *testing.T) {
	//db := GetMysqlDB2()
	//db.CreateTable(&User{})
}

type Site2 struct {
	gorm.Model
	SiteName string
}

/**
测试OpenMysqlDB请先注释掉init()
*/
func TestOpenMysqlDB2(t *testing.T) {
	//userName, password, host, port, DbName, timeOut
	//OpenMysqlDB(100, 10, "root", "123456", "127.0.0.1", "3306", "t_db", "5s")
	dbParameters := []string{
		"root",
		"123456",
		"127.0.0.1",
		"3306",
		"t_gorm2",
		"5s",
	}
	OpenMysqlDB2(100, 10, dbParameters...)
	models := []interface{}{
		&User2{},
		&Site2{},
	}
	db := GetMysqlDB2()
	_ = db.AutoMigrate(models...)
}
