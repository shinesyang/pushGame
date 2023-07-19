package model

import "time"

type YunweiData struct {
	Id               int64       `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	Platform         string      `gorm:"platform;comment:'平台号'" json:"platform"`
	PlatformNickname string      `gorm:"platform_nickname;comment:'平台别名'" json:"platform_nickname"`
	PlatformId       int         `gorm:"platform_id;comment:'平台ID'" json:"platform_id"`
	GroupName        string      `gorm:"group_name;comment:'group组名'" json:"group_name"`
	GroupId          int64       `gorm:"group_id;comment:'group组id'" json:"group_id"`
	GameIp           string      `gorm:"game_ip;comment:'内网IP'" json:"game_ip"`
	Svnurl           string      `gorm:"svnurl;comment:'group对应的svn目录'" json:"svnurl"`
	Xmx              int         `gorm:"xmx;comment:'最大内存'" json:"xmx"`
	Xms              int         `gorm:"xms;comment:'最小内存'" json:"xms"`
	Apiurl           string      `gorm:"apiurl;comment:'group配置文件调用的url'" json:"apiurl"`
	Installdir       string      `gorm:"installdir;comment:'group安装目录'" json:"installdir"`
	Logs             string      `gorm:"logs;comment:'group日志目录'" json:"logs"`
	OpenTime         time.Time   `gorm:"open_time;autoCreateTime;comment:'group配置时间'" json:"open_time"`
	LatestTime       time.Time   `gorm:"latest_time;autoUpdateTime;comment:'group更新时间'" json:"latest_time"`
	WanIp            string      `gorm:"wan_ip;comment:'外网IP'" json:"wan_ip"`
	DoName           string      `gorm:"do_name;comment:'域名'" json:"do_name"`
	GameSInfo        []GameSInfo `gorm:"ForeignKey:GroupId" json:"games_info"`
}

type GameSInfo struct {
	Id         int64  `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	GameId     int64  `gorm:"game_id;comment:'game id值'" json:"game_id"`
	GameName   string `gorm:"game_name;comment:'game名称'" json:"game_name"`
	JdbcDb     string `gorm:"jdbc_db;comment:'game连接的数据库'" json:"jdbc_db"`
	User       string `gorm:"user;comment:'数据库用户名'" json:"user"`
	Passwd     string `gorm:"passwd;comment:'数据库密码'" json:"passwd"`
	ServersStr string `gorm:"serversStr;comment:'game的ServersStr'" json:"serversStr"`
	GroupId    int64  `gorm:"group_id;comment:'game对应的group id'" json:"group_id"`
}
