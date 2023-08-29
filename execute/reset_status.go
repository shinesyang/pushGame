package execute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	"github.com/shinesyang/pushGame/overall"

	"github.com/shinesyang/pushGame/tools"

	"github.com/shinesyang/common"
)

// 重置group状态

func GetParameter(groupId, state string) map[string]string {
	timeStamp, sign := tools.CreateSign()
	parameter := make(map[string]string, 10)
	parameter["groupId"] = groupId
	parameter["state"] = state
	parameter["time"] = timeStamp
	parameter["sign"] = sign
	return parameter
}

func RestStatus(groupId, state string) {
	url := fmt.Sprintf("%s/yezimanager/common/yunwei/setServersState", overall.CParameter.Address)
	client := http.Client{
		Timeout: time.Second * 40,
	}

	parameter := GetParameter(groupId, state)

	marshal, _ := json.Marshal(parameter)

	common.Logger.Infof("打印当前参数：%v", string(marshal))

	body := bytes.NewBuffer(marshal)

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		common.Logger.Panicf("创建request实例失败:%v", err)
	}

	// 添加头部信息
	request.Header.Add(
		"User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	)

	request.Header.Add(
		"Content-Type", "application/json;charset=utf-8",
	)

	response, err := client.Do(request)

	if err != nil {
		common.Logger.Panicf("发起http请求失败: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		common.Logger.Panicf("未知的请求错误: %v", response.StatusCode)
	}

	readAll, _ := ioutil.ReadAll(response.Body)

	readAllString := string(readAll)

	//common.Logger.Infof(readAllString)

	if readAllString == "0" {
		common.Logger.Infof("重置group状态成功,当前group id为: %s", groupId)
	} else {
		common.Logger.Errorf("重置group状态失败,当前group id为: %s", groupId)
	}

}
