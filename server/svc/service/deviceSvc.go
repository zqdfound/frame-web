package userService

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"frame-web/global"
	"frame-web/svc/models"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func GetDeviceInfo(sn string) (d *models.Device, err error) {
	//info, err := reqDeviceInfo(sn)
	//if err != nil {
	//	return nil, err
	//}
	//jsonData, err := json.Marshal(info)
	//var d models.Device
	//err = json.Unmarshal(jsonData, &d)
	//if err != nil {
	//	return nil, err
	//}
	//return &d, nil
	return getDeviceResty(sn)

}

// KGF9JNHJHQ
func reqDeviceInfo(sn string) (map[string]interface{}, error) {
	// 准备Basic Auth
	username := global.CONFIG.Device.Username
	password := global.CONFIG.Device.Password
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	// 创建请求
	req, err := http.NewRequest("POST", global.CONFIG.Device.Host+"/mdm/facade/deviceInfo/"+sn, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Add("Authorization", basicAuth)
	fmt.Printf("req.Header: %v\n", req.Header)
	req.Header.Add("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()
	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
	// 解析响应
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v\n", err)
	}
	fmt.Println(result["data"].(map[string]interface{}))
	return result["data"].(map[string]interface{}), nil
}

// 使用resty工具
func getDeviceResty(sn string) (*models.Device, error) {
	// 准备Basic Auth
	username := global.CONFIG.Device.Username
	password := global.CONFIG.Device.Password
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	client := resty.New()
	defer client.SetCloseConnection(true) // 关闭连接

	res, err := resty.New().R().
		SetHeader("Authorization", basicAuth).
		SetBody(map[string]string{}). // default request content type is JSON
		SetResult(&DeviceResp{}).     // or SetResult(LoginResponse{}).
		Post(global.CONFIG.Device.Host + "/mdm/facade/deviceInfo/" + sn)

	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("请求失败，状态码: %d", res.StatusCode())
	}

	return &res.Result().(*DeviceResp).Data, nil

}

type DeviceResp struct {
	Message string        `json:"message"`
	Code    uint          `json:"code"`
	Data    models.Device `json:"data"`
}
