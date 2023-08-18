package v1

import (
	"MengCDN/internal/middleware"
	"MengCDN/internal/model"
	"MengCDN/internal/tools/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录验证
func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	_ = c.ShouldBindJSON(&data)
	code = model.CheckLogin(data.User, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.User)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
