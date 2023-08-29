package model

type GameData struct {
	Id         int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	CreateTime int32  `gorm:"column:createTime" json:"createTime"`
	State      int64  `gorm:"column:state;type:int(11)" json:"state"`
	GameId     int64  `gorm:"column:gameId" json:"game_id"`
	PlatformId int    `gorm:"column:platformId;comment:'平台ID';type:int(11)" json:"platformId"`
	BbSchema   string `gorm:"column:dbSchema;type:varchar(255)" json:"dbSchema"`
	Password   string `gorm:"column:password;type:varchar(255)" json:"password"`
	Username   string `gorm:"column:username;type:varchar(255)" json:"username"`
	//ServersStr string `gorm:"column:serversStr;type:longtext" json:"serversStr"`
	Servers    string `gorm:"column:servers;type:varchar(255)" json:"servers"`
	ServerName string `gorm:"column:serverName;type:varchar(255)" json:"server_name"`
	MixServer  string `gorm:"column:mixServer;type:varchar(255)" json:"mixServer"`
}

func (GameData) TableName() string {
	return "gamedata"
}
