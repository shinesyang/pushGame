package main

import (
	"flag"
	"strings"

	"gitee.com/shinesyang/pushGame/execute"

	"github.com/shinesyang/common"
)

// 重置接口,人工处理

func main() {
	groupId := flag.String("groupId", "", "groupId 为group的id,多个ID使用,分隔")
	state := flag.String("state", "", `
	重置group接口状态,状态类型:
	1 - 正常
	2 - 维护
	3 - 停服
	4 - 保持版本
`)
	flag.Parse()

	if *groupId == "" || *state == "" {
		common.Logger.Panic("groupId为必传参数")
	}
	groupIds := strings.Split(*groupId, ",")

	for _, j := range groupIds {
		execute.RestStatus(j, *state)
	}

}
