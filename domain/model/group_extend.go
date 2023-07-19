package model

import (
	"time"
)

// 扩展group字段

type GroupExtend struct {
	Id           int64     `gorm:"id;primary_key;not_null;auto_increment" json:"id"`
	Svnurl       string    `gorm:"svnurl;comment:'group对应的svn目录';type:varchar(255)" json:"svnurl"`
	Xmx          int       `gorm:"xmx;comment:'最大内存'" json:"xmx"`
	Xms          int       `gorm:"xms;comment:'最小内存'" json:"xms"`
	Apiurl       string    `gorm:"apiurl;comment:'group配置文件调用的url';type:varchar(255)" json:"apiurl"`
	Installdir   string    `gorm:"installdir;comment:'group安装目录';type:varchar(255)" json:"installdir"`
	Logs         string    `gorm:"logs;comment:'group日志目录';type:varchar(255)" json:"logs"`
	OpenTime     time.Time `gorm:"open_time;autoCreateTime;comment:'group配置时间'" json:"open_time"`
	LatestTime   time.Time `gorm:"latest_time;autoUpdateTime;comment:'group更新时间'" json:"latest_time"`
	GameIp       string    `gorm:"game_ip;comment:'内网IP';type:varchar(255)" json:"game_ip"`
	WanIp        string    `gorm:"wan_ip;comment:'外网IP';type:varchar(255)" json:"wan_ip"`
	DoName       string    `gorm:"do_name;comment:'域名';type:varchar(255)" json:"do_name"`
	GroupId      int64     `gorm:"group_id;comment:'group组id'" json:"group_id"`
	GroupVersion int       `gorm:"group_version;type int(11);comment:'group 版本'" json:"group_version"`
	DataVersion  int       `gorm:"data_version;type int(11);comment:'data 版本'" json:"data_version"`
}

// 设置表名
func (GroupExtend) TableName() string {
	return "groupextend"
}

// 查询版本参数使用
type Version struct {
	DataVersion  int `json:"data_version"`
	GroupVersion int `json:"group_version"`
}

// 查询group 地址信息使用
type GroupAddress struct {
	GameIp  string `json:"game_ip"`
	WanIp   string `json:"wan_ip"`
	DoName  string `json:"do_name"`
	CountId int32  `json:"count_id"`
}
