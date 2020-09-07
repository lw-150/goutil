package goutil

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//mysql dsn格式
//涉及参数:
//username   数据库账号
//password   数据库密码
//host       数据库连接地址，可以是Ip或者域名
//port       数据库端口
//Dbname     数据库名
//  charset=utf8 客户端字符集为utf8
//  parseTime=true 支持把数据库datetime和date类型转换为golang的time.Time类型
//  loc=Local 使用系统本地时区
//gorm 设置mysql连接超时参数
//开发的时候经常需要设置数据库连接超时参数，gorm是通过dsn的timeout参数配置
//例如，设置10秒后连接超时，timeout=10s
//设置读写超时时间
// readTimeout - 读超时时间，0代表不限制
// writeTimeout - 写超时时间，0代表不限制
//root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s
var _db *gorm.DB

//func init() {
//	var err error
//	//dsn:="root:123456@tcp(localhost:3306)/tizi365?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s"
//	userName := "root"   //账号
//	password := "123456" //密码
//	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
//	port := "3306"       //数据库端口
//	DbName := "t_db"     //数据库名
//	timeOut := "3s"      //连接超时，3秒
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", userName, password, host, port, DbName, timeOut)
//	_db, err = gorm.Open("mysql", dsn)
//	if err != nil {
//		panic("连接数据库失败, error=" + err.Error())
//	}
//	//设置了连接池参数，则默认使用数据库连接池，用连接池，则不能db.close()
//	_db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
//	_db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
//	_db.SingularTable(true)
//}

func OpenMysqlDB(maxOpenConn int, maxIdleConn int, dbParameters ...string) {
	var err error
	userName := dbParameters[0] //账号
	password := dbParameters[1] //密码
	host := dbParameters[2]     //数据库地址，可以是Ip或者域名
	port := dbParameters[3]     //数据库端口
	DbName := dbParameters[4]   //数据库名
	timeOut := dbParameters[5]  //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", userName, password, host, port, DbName, timeOut)
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//设置了连接池参数，则默认使用数据库连接池；用连接池，则不能db.close()否则会导致连接池关闭；用连接池，则多读写情况下需用事务
	_db.DB().SetMaxOpenConns(maxOpenConn) //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(maxIdleConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	//true:表名单数
	_db.SingularTable(true)
	// 设置表结构的存储引擎为InnoDB
	_db.Set("gorm:table_options", "ENGINE=InnoDB")
}
func GetMysqlDB() *gorm.DB {
	return _db
}
