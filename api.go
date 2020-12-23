package safecustody_sdk_go

import (
	"crypto/md5"
	"fmt"
	"time"
)

//api
type Api struct {
	Host      string
	user      UserInfo
	timestamp int64
}

//用户信息
type UserInfo struct {
	appId     string
	secretKey string
	userId    string
	apiKey    string
}

//auth
type auth struct {
	Token     string `json:"token"`
	Timestamp int64  `json:"timestamp"`
	ApiKey    string `json:"api_key"`
}

//用来发送给服务器的请求参数
type param struct {
	Data     interface{} `json:"data"`
	AppId    string      `json:"appid"`
	Cryptype int         `json:"cryptype"`
}

//设置用户信息
func (a *Api) SetUserInfo(appId, secretKey, userId, apiKey string) {
	a.user.appId = appId
	a.user.secretKey = secretKey
	a.user.userId = userId
	a.user.apiKey = apiKey
}

//设置token
func (a *Api) setToken(s string) string {
	str := a.user.apiKey + "_" + a.user.secretKey + "_" + a.user.userId + "_" + s
	return Md5(str)
}

//获取用户验证
func (a *Api) getAuth() auth {
	a.timestamp = time.Now().Unix()
	s := fmt.Sprint(time.Now().Unix())
	return auth{
		Token:     a.setToken(s),
		Timestamp: a.timestamp,
		ApiKey:    a.user.apiKey,
	}
}

//组建用来发送给服务器的请求参数
func (a *Api) buildParam(data interface{}) param {
	return param{
		Data:  data,
		AppId: a.user.appId,
	}
}

//提币的签名
func (a *Api) WithdrawSign(addr, memo, usertags string) string {
	s := a.user.apiKey + "_" + a.user.secretKey + "_" + a.user.userId + "_" + fmt.Sprint(a.timestamp) + "_" + addr + "_" + memo + "_" + usertags
	return Md5(s)
}

func Md5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
