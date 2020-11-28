package goutil

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
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
var _db2 *gorm.DB

//https://www.kancloud.cn/sliver_horn/gorm/1861152
//以下基于gorm2
func OpenMysqlDB2(maxOpenConn int, maxIdleConn int, dbParameters ...string) {
	var err error
	userName := dbParameters[0] //账号
	password := dbParameters[1] //密码
	host := dbParameters[2]     //数据库地址，可以是Ip或者域名
	port := dbParameters[3]     //数据库端口
	DbName := dbParameters[4]   //数据库名
	timeOut := dbParameters[5]  //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", userName, password, host, port, DbName, timeOut)
	_db2, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//true:表名单数
	//_db2.SingularTable(true)
	// 设置表结构的存储引擎为InnoDB
	_db2.Set("gorm:table_options", "ENGINE=InnoDB")
	//设置了连接池参数，则默认使用数据库连接池；用连接池，则不能db.close()否则会导致连接池关闭；用连接池，则多读写情况下需用事务
	sqlDB, err := _db2.DB()
	if err == nil {
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(maxIdleConn)
		// SetMaxOpenConns 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(maxOpenConn)
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Second * 4)
	}
}
func GetMysqlDB2() *gorm.DB {
	return _db2
}
