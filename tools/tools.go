package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// md5加密
func CreateSign() (string, string) {
	now := time.Now().UnixMilli()
	nowStr := strconv.FormatInt(now, 10)
	value := nowStr + "yunwei_qb_cqhy"
	d := []byte(value)
	m := md5.New()
	m.Write(d)
	return nowStr, hex.EncodeToString(m.Sum(nil))
}
