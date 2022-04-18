package global

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/giant-stone/go/ghttp"
	"github.com/giant-stone/go/gutil"

	"github.com/WuhuGoldPilot/SecondKill/logs"
)

func NotifyUser(v ...interface{}) {
	if PushToken == "" {
		logs.PrintlnInfo("没有推送token，推送失败！")
		return
	}

	chunks := []string{}
	for _, chunk := range v {
		chunks = append(chunks, fmt.Sprintf("%v", chunk))
	}
	msg := strings.Join(chunks, " ")

	fullurl := PushplusReqURL
	reqBody := PushplusReq{
		Token:   PushToken,
		Title:   "秒杀程序通知",
		Content: msg,
		Channel: "wechat",
	}
	rqBody, _ := json.Marshal(&reqBody)
	rq := ghttp.New().
		SetTimeout(time.Second * 3).
		SetRequestMethod("POST").
		SetUri(fullurl).
		SetPostBody(&rqBody)
	err := rq.Send()
	if gutil.CheckErr(err) {
		logs.PrintErr("notify user fail", string(rq.RespBody))
	}
}
