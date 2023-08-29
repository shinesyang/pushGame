package model

// 服务器配置

type ServerData struct {
	Id     int64  `gorm:"column:id;type:bigint(20);primary_key;not_null;auto_increment" json:"id"`
	WanIp  string `gorm:"column:wanIp;type:varchar(255)" json:"wanIp"`
	GameIp string `gorm:"column:gameIp;type:varchar(255)" json:"gameIp"`
	DoName string `gorm:"column:doName;type:varchar(255)" json:"doName"`
	//SshPort int    `gorm:"column:sshPort;type:int(11)" json:"sshPort"`
	State     int  `gorm:"column:state;type:int(11)" json:"state"`
	BigVolume bool `gorm:"column:bigVolume;type:bool" json:"bigVolume"`
	Recycle   bool `gorm:"column:recycle;type:bool;comment:'是否回收'" json:"recycle"`
	Num       int  `gorm:"-:all" json:"num"` // 这个字段不生成到数据库
}

func (ServerData) TableName() string {
	return "serverdata"
}
