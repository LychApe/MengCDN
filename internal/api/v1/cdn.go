package v1

import (
	"MengCDN/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CdnSW 模块开关信息
func CdnSW(c *gin.Context) {
	id := c.Param("id")
	mk := c.Param("mk")
	code := model.CdnSW(id, mk)
	if code == "1" {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "开启",
		})
	} else if code == "0" {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "关闭",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "查询异常",
		})
	}
}

// PUTCdnSW 编辑模块开关信息
func PUTCdnSW(c *gin.Context) {
	id := c.Param("id")
	mk := c.Param("mk")
	sw := c.Param("sw")
	code := model.PUTCdnSW(id, mk, sw)
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "修改成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "修改失败！",
		})
	}
}

// CdnWL 模块白名单信息
func CdnWL(c *gin.Context) {
	mk := c.Param("mk")
	data := model.CdnWL(mk)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   strings.ReplaceAll(data, "\"", ""),
	})
}

// PUTCdnWL 编辑模块白名单信息
func PUTCdnWL(c *gin.Context) {
	var data model.WhiteList
	mk := c.Param("mk")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		println("ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo")
	}
	code := model.PUTCdnWL(mk, &data)
	if code == 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    data.Value,
			"message": "修改成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "修改失败！",
		})
	}
}
