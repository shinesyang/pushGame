package model

type LinkData struct {
	Id            int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	Host          string `gorm:"host" json:"host"`
	Port          int    `gorm:"port" json:"port"`
	GmPort        int    `gorm:"column:gmPort;type:int(11)" json:"gmPort"`
	WebsocketPort int    `gorm:"column:websocketPort;type:int(11)" json:"websocketPort"`
	DoName        string `gorm:"column:doName;type:varchar(255)" json:"doName"`
}

func (LinkData) TableName() string {
	return "linkdata"
}
