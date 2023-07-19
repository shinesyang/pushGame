package model

type GroupData struct {
	Id               int64       `gorm:"column:id;primary_key;not_null;auto_increment" json:"id"`
	CreateTime       int32       `gorm:"column:createTime" json:"createTime"`
	GameIds          string      `gorm:"column:gameIds;type:varchar(255)" json:"gameIds"`
	Link             int64       `gorm:"column:link;type:int(11)" json:"link"`
	LogPath          string      `gorm:"column:logPath;type:varchar(255)" json:"logPath"`
	Proxy            int64       `gorm:"column:proxy;type:int(11)" json:"proxy"`
	Ready            int64       `gorm:"column:ready;type:tinyint(4)" json:"ready"`
	Report           int64       `gorm:"column:report;type:int(11)" json:"report"`
	ShareHost        string      `gorm:"column:shareHost;type:varchar(255)" json:"shareHost"`
	SharePort        int         `gorm:"column:sharePort;type:int(11)" json:"sharePort"`
	FightIds         string      `gorm:"column:fightIds;type:varchar(255)" json:"fightIds"`
	WorldIds         string      `gorm:"column:worldIds;type:varchar(255)" json:"worldIds"`
	GroupDesc        string      `gorm:"column:groupDesc;type:longtext" json:"groupDesc"`
	Platform         string      `gorm:"column:platform;type:varchar(255);comment:'平台号'" json:"platform"`
	PlatformNickname string      `gorm:"column:platformNickname;type:varchar(255);comment:'平台别名'" json:"platformNickname"`
	PlatformId       int         `gorm:"column:platformId;comment:'平台ID';type:int(11)" json:"platformId"`
	GroupExtend      GroupExtend `gorm:"ForeignKey:GroupId" json:"group_extend"`
}

// 设置表名
func (GroupData) TableName() string {
	return "groupdata"
}
