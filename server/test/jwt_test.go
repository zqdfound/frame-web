package test

import (
	"encoding/json"
	"frame-web/middleware"
	"testing"
)

func Test_decodeJwt(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAzMDE1OTQsInVzZXJfaWQiOjEyMywidXNlcm5hbWUiOiJ0ZXN0dXNlciJ9.BVCBOGDCp5Ribgfe7S4SmSSTxLjN_o9vBx-7GhldKQg"
	parseToken, err := middleware.ParseToken(token, "woailiming")
	if err != nil {
		t.Errorf("ParseToken error: %v", err)
		return
	}

	// 打印解析后的token内容
	t.Logf("Parsed token: %+v", parseToken)

	// 或者格式化输出
	jsonData, _ := json.MarshalIndent(parseToken, "", "  ")
	t.Logf("Formatted token claims:\n%s", string(jsonData))
}
