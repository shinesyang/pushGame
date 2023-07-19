package model

type FightData struct {
	Id         int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	CreateTime int32  `gorm:"column:createTime;type:(bigint(20))" json:"createTime"`
	DbSchema   string `gorm:"column:dbSchema;type:varchar(255)" json:"dbSchema"`
	Password   string `gorm:"column:password;type:varchar(255)" json:"password"`
	State      int64  `gorm:"column:state;type:tinyint(4)" json:"state"`
	Username   string `gorm:"column:username;type:varchar(255)" json:"username"`
	Host       string `gorm:"column:host;type:varchar(255)" json:"host"`
	Port       int    `gorm:"column:port;type:int(11)" json:"port"`
}
