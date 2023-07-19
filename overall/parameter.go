package overall

import (
	"os"
	"path"

	"github.com/shinesyang/common"

	"github.com/BurntSushi/toml"
)

type ConfigParameter struct {
	YunweiDB         string `json:"yunweiDB" toml:"yunweiDB"`
	YunweiPasswd     string `json:"yunweiPasswd" toml:"yunweiPasswd"`
	YunweiUserName   string `json:"yunweiUserName" toml:"yunweiUserName"`
	YunweiServerPort string `json:"yunweiServerPort" toml:"yunweiServerPort"`
	YunweiHost       string `json:"yunweiHost" toml:"yunweiHost"`
	Address          string `json:"address" toml:"address"`
}

var CParameter ConfigParameter

// 解析配置文件
func ReadConfig() {
	dir, _ := os.Getwd()
	CParameter = ConfigParameter{}

	fileName := path.Join(dir, "config", "config")
	if _, err := toml.DecodeFile(fileName, &CParameter); err != nil {
		common.Logger.Panicf("解析配置文件失败: %v", err)
	}
}

func init() {
	ReadConfig()
}
