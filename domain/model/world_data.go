package model

type WorldData struct {
	Id         int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	CreateTime int32  `gorm:"createTime" json:"createTime"`
	DbSchema   string `gorm:"dbSchema" json:"dbSchema"`
	Password   string `gorm:"password" json:"password"`
	State      int64  `gorm:"state" json:"state"`
	Username   string `gorm:"username" json:"username"`
}
