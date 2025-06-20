package middleware

import (
	"encoding/json"
	"errors"
	"frame-web/model/response"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserContext struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Exp      uint64 `json:"exp"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SigningKey  string   // 签名密钥
	WhiteList   []string // 白名单路径
	ContextKey  string   // 上下文键名
	TokenLookup string   // token查找方式，如 "header:Authorization"
}

func JWTAuth2() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.FailWithMessage("测试需要鉴权", c)
		c.Abort()
		return
	}
}

// JWTAuth JWT鉴权中间件
func JWTAuth(config JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查白名单
		for _, path := range config.WhiteList {
			// 支持通配符*的路径匹配
			if match, _ := filepath.Match(path, c.Request.URL.Path); match {
				c.Next()
				return
			}

			// 保留原有的前缀匹配
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next()
				return
			}
		}

		// 获取token
		token, err := getToken(c, config.TokenLookup)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// 解析token
		claims, err := ParseToken(token, config.SigningKey)
		if err != nil {
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			response.FailWithMessage("无效的token", c)
			return
		}

		// 设置用户上下文
		c.Set(config.ContextKey, claims)
		c.Next()
	}
}

// getToken 从请求中获取token
func getToken(c *gin.Context, lookup string) (string, error) {
	parts := strings.Split(lookup, ":")
	if len(parts) != 2 {
		return "", errors.New("无效的token查找配置")
	}

	switch parts[0] {
	case "header":
		return c.GetHeader(parts[1]), nil
	case "query":
		return c.Query(parts[1]), nil
	// case "cookie":
	// 	return c.GetCookie(parts[1])
	default:
		return "", errors.New("不支持的token查找方式")
	}
}

// ParseToken 解析JWT token
func ParseToken(tokenString, signingKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

func ParseToken2User(tokenString string, jwtConfig *JWTConfig) (*UserContext, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		}
		return []byte(jwtConfig.SigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userContextMap, ok := claims[jwtConfig.ContextKey].(map[string]interface{})
		if !ok {
			return nil, errors.New("token中的usercontext格式不正确")
		}
		userContextJSON, err := json.Marshal(userContextMap)
		if err != nil {
			return nil, err
		}
		var userContext UserContext
		json.Unmarshal(userContextJSON, &userContext)
		return &userContext, nil
	}

	return nil, errors.New("无效的token")
}

// GenerateToken 根据用户信息生成JWT token
func GenerateToken(signingKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func GenerateTokenByUser(jwtConfig *JWTConfig, user *UserContext) (string, error) {
	claims := jwt.MapClaims{
		jwtConfig.ContextKey: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.SigningKey))
}

// // 生成token示例
// claims := jwt.MapClaims{
// 	"user_id": 123,
// 	"username": "testuser",
// 	"exp": time.Now().Add(time.Hour * 24).Unix(), // 24小时后过期
// }

// token, err := middleware.GenerateToken("your-secret-key", claims)
// if err != nil {
// 	// 处理错误
// }
