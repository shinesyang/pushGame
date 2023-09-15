package execute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/shinesyang/pushGame/overall"
	"github.com/shinesyang/pushGame/result"

	"github.com/shinesyang/pushGame/tools"

	"gorm.io/gorm"

	"github.com/shinesyang/pushGame/domain/model"

	"github.com/shinesyang/common"
)

/*
	推送新增的game到管理后台
	managerUrl/common/yunwei/yunweiAddServers
	接口地址
*/

type GroupGame struct {
	ConnDB         *gorm.DB
	GroupDataModel model.GroupData
	GameDataModel  []model.GameData
	linkDataModel  model.LinkData
}

func NewGroupGame(db *gorm.DB) *GroupGame {
	return &GroupGame{
		ConnDB: db,
	}
}

// 添加group and game
func (e *GroupGame) Push() {
	e.GetData()
	e.PushGameToManager()

}

// 获取group(数据库)
func (e *GroupGame) GetGroup(id int64) {
	// 查询group
	groupDataModel := model.GroupData{}
	if err := e.ConnDB.Model(model.GroupData{}).Preload("GroupExtend").Where("id = ?", id).
		First(&groupDataModel).Error; err != nil {
		common.Logger.Panicf("查询groupdata出错: %v", err)
	}
	//
	e.GroupDataModel = model.GroupData{}
	e.GroupDataModel = groupDataModel
}

// 传入group
func (e *GroupGame) WithGroup(groupData model.GroupData) {
	e.GroupDataModel = model.GroupData{}
	e.GroupDataModel = groupData
}

func (e *GroupGame) GetData() {
	groupDataModel := e.GroupDataModel
	// 获取查询参数
	platformId := groupDataModel.PlatformId
	gameIds := groupDataModel.GameIds
	gameIdsList := strings.Split(gameIds, ",")
	link := groupDataModel.Link

	//e.GroupDataModel = groupDataModel

	// 查询game
	gameDataModel := []model.GameData{}
	e.GameDataModel = []model.GameData{}
	if len(gameIdsList) > 0 {
		if err := e.ConnDB.Model(model.GameData{}).Where("platformId = ? AND gameId IN ?", platformId, gameIdsList).
			Find(&gameDataModel).Error; err != nil {
			common.Logger.Panicf("查询gamedata出错: %v", err)
		}
		e.GameDataModel = gameDataModel
	}

	// 查询link
	linkDataModel := model.LinkData{}
	e.linkDataModel = model.LinkData{}
	if link != 0 {
		if err := e.ConnDB.Model(model.LinkData{}).Where("id = ?", link).First(&linkDataModel).Error; err != nil {
			common.Logger.Panicf("查询linkdata出错: %v", err)
		}
		e.linkDataModel = linkDataModel
	}

}

func (e *GroupGame) GetParam() map[string]interface{} {
	Parameter := make(map[string]interface{}, 20)
	Game := []map[string]string{}

	// group
	timeStamp, sign := tools.CreateSign()
	Parameter["sign"] = sign
	Parameter["time"] = timeStamp
	Parameter["groupId"] = e.GroupDataModel.Id
	Parameter["proxy"] = e.GroupDataModel.Proxy
	Parameter["ready"] = e.GroupDataModel.Ready
	Parameter["report"] = e.GroupDataModel.Report
	Parameter["shareHost"] = e.GroupDataModel.ShareHost
	Parameter["sharePort"] = e.GroupDataModel.SharePort
	Parameter["fightIds"] = e.GroupDataModel.FightIds
	Parameter["worldIds"] = e.GroupDataModel.WorldIds
	Parameter["groupDesc"] = e.GroupDataModel.GroupDesc
	Parameter["platformId"] = e.GroupDataModel.PlatformId
	Parameter["groupVersion"] = e.GroupDataModel.GroupExtend.GroupVersion
	Parameter["dataVersion"] = e.GroupDataModel.GroupExtend.DataVersion
	Parameter["groupIp"] = e.GroupDataModel.GroupExtend.GameIp
	Parameter["platformNickname"] = e.GroupDataModel.PlatformNickname
	Parameter["logPath"] = e.GroupDataModel.LogPath

	// game
	for _, game := range e.GameDataModel {
		gameMap := make(map[string]string, 15)
		gameId := game.GameId
		gameIdStr := strconv.FormatInt(gameId, 10)
		gameMap["serverId"] = gameIdStr
		gameMap["serverName"] = game.ServerName
		gameMap["mixServer"] = game.MixServer
		gameMap["servers"] = game.Servers
		Game = append(Game, gameMap)
	}
	Parameter["gameData"] = Game

	// link
	Parameter["serverIp"] = e.linkDataModel.Host
	Parameter["serverDomain"] = e.linkDataModel.DoName
	Parameter["httpPort"] = e.linkDataModel.GmPort
	Parameter["serverPort"] = e.linkDataModel.WebsocketPort

	return Parameter
}

func (e *GroupGame) PushGameToManager() {
	url := fmt.Sprintf("%s/yezimanager/common/yunwei/yunweiAddServers", overall.CParameter.Address)
	client := http.Client{
		Timeout: time.Second * 40,
	}

	// 添加参数
	param := e.GetParam()

	marshal, _ := json.Marshal(param)

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

	readAll, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		common.Logger.Errorf("输出返回结果: %v", string(readAll))
		common.Logger.Panicf("未知的请求错误: %v,当前code: %d", err, response.StatusCode)
	}

	//readAllString := string(readAll)

	// 获取返回参数
	resultData := result.ResultData{}
	if err1 := json.Unmarshal(readAll, &resultData); err1 != nil {
		common.Logger.Panicf("请求返回参数解析失败: %v", err1)
	}

	marshal1, _ := json.Marshal(resultData)
	common.Logger.Debugf("返回请求结果: %s", string(marshal1))

	if resultData.Code == 0 {
		succServers := resultData.SuccServers
		failServers := resultData.FailServers
		common.Logger.Infof("添加成功,当前group id为: %d,group名: %s", e.GroupDataModel.Id, e.GroupDataModel.GroupDesc)
		if len(succServers) > 0 {
			common.Logger.Infof("添加成功的game有: %s", strings.Join(succServers, ","))
		}
		if len(failServers) > 0 {
			common.Logger.Errorf("添加失败的的game有: %s", strings.Join(failServers, ","))
		}

	} else {
		common.Logger.Errorf("当前group id为: %d,group名: %s, 添加失败: %s",
			e.GroupDataModel.Id, e.GroupDataModel.GroupDesc, resultData.Msg,
		)
	}

}
