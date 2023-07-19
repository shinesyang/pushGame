package model

type ReportData struct {
	Id   int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	Host string `gorm:"host" json:"host"`
	Port int    `gorm:"port" json:"port"`
}
