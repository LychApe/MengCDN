package middleware

import (
	"MengCDN/internal/tools/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte("2DH5jdjh267")

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// SetToken 生成Token
func SetToken(username string) (string, int) {
	RegisteredClaims := MyClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "MengCDN",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, RegisteredClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	} else {
		return token, errmsg.SUCCESS
	}
}

// CheckToken 验证Token
func CheckToken(token string) (*MyClaims, int) {
	regToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := regToken.Claims.(*MyClaims); regToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtToken JWT中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ErrorTokeExist
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokeTypeWrong
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tokenCode := CheckToken(checkToken[1])
		if tokenCode == errmsg.ERROR {
			code = errmsg.ErrorTokeWrong
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if jwt.NewNumericDate(time.Now()).Unix() > key.ExpiresAt.Unix() {
			code = errmsg.ErrorTokeRuntime
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}

}
