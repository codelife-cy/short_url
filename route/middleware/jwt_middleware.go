package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "car_management"
)

func GetSignKey() string {
	return SignKey
}

func SetSignKey(signKey string) string {
	SignKey = signKey
	return SignKey
}

// 获取一个JWT 对象

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(GetSignKey())}
}

// 自定义结构体参数

type CustomClaims struct {
	ID       uint   `form:"id"`
	Phone    string `form:"phone"`
	Type     int    `form:"type"`
	RealName string `form:"real_name"`
	jwt.StandardClaims
}

// 生成token

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.SigningKey)
}

// 授权验证

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		split := strings.Split(path, "/")
		method := split[len(split)-1]
		if method == "login" {
			context.Next()
			return
		}
		token := context.GetHeader("Authorization")
		if len(token) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, "请求未授权")
			return
		}
		newJWT := NewJWT()
		claims, err := newJWT.ParseToken(token)
		if err != nil {
			// token过期
			if err == TokenExpired {
				if token, err = newJWT.RefreshToken(token); err == nil {
					context.Abort()
					context.JSON(http.StatusUnauthorized, "授权已过期")
					return
				}
			}
		}
		context.Set("claims", claims)
		context.Next()
		return
	}
}

// 解析token

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token 过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	// 获取解析token中的claims(CustomClaims)
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 刷新token

func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
