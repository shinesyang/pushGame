package utils

// 定义各种struct类型参数等等

//数据库连接参数
type DBConnArgs struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Db     string `json:"db"`
	Host   string `json:"host"`
	Port   string `json:"port"`
}
