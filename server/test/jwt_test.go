package test

import (
	"frame-web/middleware"
	"testing"
)

// func Test_decodeJwt(t *testing.T) {
// 	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyY29udGV4dCI6eyJ1c2VyX2lkIjoiMTIzIiwidXNlcm5hbWUiOiJ0ZXN0dXNlciIsImV4cCI6MTc1MDMxNjc5OX19.PXX4NltLjIr6RGT01X28HmeUkYgVQ6O7fshYL881HRE"
// 	parseToken, err := middleware.ParseToken(token, "woailiming")
// 	if err != nil {
// 		t.Errorf("ParseToken error: %v", err)
// 		return
// 	}

// 	// 打印解析后的token内容
// 	t.Logf("Parsed token: %+v", parseToken)

// 	// 或者格式化输出
// 	jsonData, _ := json.MarshalIndent(parseToken, "", "  ")
// 	t.Logf("Formatted token claims:\n%s", string(jsonData))
// }

func Test_decodeUser(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyY29udGV4dCI6eyJ1c2VyX2lkIjoiMTIzIiwidXNlcm5hbWUiOiJ0ZXN0dXNlciIsImV4cCI6MTc1MDMxNjc5OX19.PXX4NltLjIr6RGT01X28HmeUkYgVQ6O7fshYL881HRE"
	jwtConfig := middleware.JWTConfig{
		SigningKey: "woailiming",
		// WhiteList:   []string{"/api/public", "/test/set/jwt"},
		ContextKey:  "usercontext",
		TokenLookup: "header:Authorization",
	}

	userContext, err := middleware.ParseToken2User(token, &jwtConfig)
	if err != nil {
		t.Errorf("ParseToken2User error: %v", err)
		return
	}

	// 打印返回的UserContext实体
	t.Logf("UserContext: %+v", userContext)
}
