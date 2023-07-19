package utils

import (
	"github.com/shinesyang/pushGame/overall"

	"github.com/shinesyang/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 运维数据库连接
func ConnDB() *gorm.DB {
	yuDbArgs := DBConnArgs{
		User:   overall.CParameter.YunweiUserName,
		Passwd: overall.CParameter.YunweiPasswd,
		Db:     overall.CParameter.YunweiDB,
		Host:   overall.CParameter.YunweiHost,
		Port:   overall.CParameter.YunweiServerPort,
	}

	// 连接mysql v1版本 (v1.20.0以下版本)
	//db, err := gorm.Open("mysql", yuDbArgs.User+":"+yuDbArgs.Passwd+"@("+yuDbArgs.Host+":"+yuDbArgs.Port+")/"+yuDbArgs.Db+"?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	common.Logger.Panicf("连接mysql数据库失败: %v", err)
	//}
	//db.SingularTable(true) // 禁止复数表

	// 连接mysql v2版本(v1.20.0以上版本)
	dsn := yuDbArgs.User + ":" + yuDbArgs.Passwd + "@(" + yuDbArgs.Host + ":" + yuDbArgs.Port + ")/" + yuDbArgs.Db + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止复数表
		},
	})

	if err != nil {
		common.Logger.Panicf("连接mysql数据库失败: %v", err)
	}

	return db
}
