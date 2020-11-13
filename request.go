package safecustody_sdk_go

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

//服务端响应的包体
type respBody struct {
	Cryptype int `json:"cryptype"`
	Data     struct {
		Eno  int         `json:"eno"`
		Emsg string      `json:"emsg"`
		Data interface{} `json:"data"`
	} `json:"data"`
}

//post请求
func Post(url string, data param) ([]byte, error) {
	var err error
	var resp *http.Response

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err = client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}

//这里用来发送请求
func (a *Api) request(url string, d param, dataTypes interface{}) error {
	l := len(a.Host)
	if l <= 0 {
		return sdkError("host can not nil")
	}
	if a.Host[l-1:] != "/" {
		a.Host = a.Host + "/"
	}
	r := &respBody{}
	res, err := Post(a.Host+url, d)
	if err != nil {
		return sdkError(err.Error())
	}

	if err = json.Unmarshal(res, r); err != nil {
		return sdkError(err.Error())
	}

	//接口返回的错误
	if r.Data.Eno != 0 {
		return errors.New(r.Data.Emsg)
	}
	if dataTypes != nil {
		err = unmarshal(r.Data.Data, &dataTypes)
	}
	return sdkError(err.Error())
}

func unmarshal(data interface{}, dataType interface{}) (err error) {
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &dataType)
	if err != nil {
		return
	}
	return
}

func sdkError(err string) error {
	return errors.New("[SDK ERR]:" + err)
}
