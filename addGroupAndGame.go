package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/shinesyang/pushGame/execute"

	"github.com/shinesyang/pushGame/domain/utils"

	"github.com/shinesyang/common"
)

// 后台推送,人工处理
var Id string

func init() {
	groupId := flag.String("group_id", "", "推送到后台的group_id,多个ID使用,分隔")
	flag.Parse()
	Id = *groupId
}

func main() {
	if Id == "" {
		common.Logger.Panic("group_id为必传参数")
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
