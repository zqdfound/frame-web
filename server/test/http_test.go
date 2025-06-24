package test

import (
	"encoding/base64"
	"fmt"
	"frame-web/svc/models"
	"testing"

	"github.com/go-resty/resty/v2"
)

func Test_req(t *testing.T) {
	auth := ""
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	client := resty.New()
	defer client.SetCloseConnection(true) // 关闭连接
	res, err := client.R().
		SetHeader("Authorization", basicAuth).
		SetBody(map[string]string{}). // default request content type is JSON
		SetResult(&DeviceResp{}).     // or SetResult(LoginResponse{}).
		//SetError(&LoginError{}).      // or SetError(LoginError{}).
		Post("")

	fmt.Println("===================")
	//fmt.Println(err, res)
	fmt.Println(err, res)
	fmt.Println("===================")
	fmt.Println(res.Result().(*DeviceResp).Data.SerialNo)
	//fmt.Println(res.Error())
	fmt.Println("===================")
}

type DeviceResp struct {
	Message string        `json:"message"`
	Code    uint          `json:"code"`
	Data    models.Device `json:"data"`
}
