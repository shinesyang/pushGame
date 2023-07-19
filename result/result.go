package result

// 接口请求返回参数
type ResultData struct {
	Code        int      `json:"code"`
	Msg         string   `json:"msg"`
	SuccServers []string `json:"succServers"`
	FailServers []string `json:"failServers"`
}
