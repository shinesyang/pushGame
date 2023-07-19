package main

import (
	"flag"
	"strconv"
	"strings"

	"gitee.com/shinesyang/pushGame/execute"

	"gitee.com/shinesyang/pushGame/domain/utils"

	"github.com/shinesyang/common"
)

// 后台推送,人工处理
var Id string

func init() {
	platformId := flag.String("platform_id", "", "推送到后台的platform_id,多个ID使用,分隔")
	flag.Parse()
	Id = *platformId
}

func main() {
	if Id == "" {
		common.Logger.Panic("platform_id为必传参数")
	}
	Ids := strings.Split(Id, ",")
	connDB := utils.ConnDB()
	groupGame := execute.NewGroupGame(connDB)
	for _, i := range Ids {
		atoi, _ := strconv.Atoi(i)
		groupGame.GetGroup(int64(atoi))
		groupGame.Push()
	}
}