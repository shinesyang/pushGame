package model

type PlatformData struct {
	Id          int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	Name        string `gorm:"name" json:"name"`
	ShareHost   string `gorm:"shareHost" json:"shareHost"`
	SharePort   int    `gorm:"sharePort" json:"sharePort"`
	UrlRecharge string `gorm:"urlRecharge" json:"urlRecharge"`
}
