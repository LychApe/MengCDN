package service

import (
	"MengCDN/internal/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func WpProxyTh(c *gin.Context) {
	if GetCDNSWC("WPth") == "on" {
		UserIP := c.ClientIP()
		wpPackage := c.Param("package")
		wpVersion := c.Param("version")
		wpPath := c.Param("path")
		//eg1:https://themes.svn.wordpress.org/2012-underscores/1.4/
		target := "themes.svn.wordpress.org"
		wpProxyPath := "/" + wpPackage + "/" + wpVersion + "/" + wpPath

		var whiteListSW = GetWLSWC("WPth")

		if whiteListSW == "on" {
			// 白名单检测
			var authBool bool = false
			for _, i := range GetWLC("WPthWL") {
				if i == wpPackage {
					// 状态码检测
					if tools.CheckHttpCode1(target+wpProxyPath) == "SM00002" {
						c.HTML(http.StatusOK, "error.html", gin.H{
							"title":      "未寻找到文件",
							"statusCode": "404",
							"message1":   "未寻找到文件",
							"message2":   "在这个WP仓库中未寻找到这个文件。",
							"message3":   "请检查你的路径信息。",
							"userIp":     UserIP,
						})
					} else if tools.CheckHttpCode1(target+wpProxyPath) == "SM00001" {
						c.HTML(http.StatusOK, "error.html", gin.H{
							"title":      "代理模块异常",
							"statusCode": "M100002",
							"message1":   "代理模块异常",
							"message2":   "服务器与互联网连接异常！",
							"message3":   "联系管理员处理。",
							"userIp":     UserIP,
						})
					} else {
						// 反向代理模块
						director := func(req *http.Request) {
							req.URL.Scheme = "https"
							req.URL.Host = target
							req.URL.Path = wpProxyPath
							req.Host = target
						}
						proxy := httputil.ReverseProxy{Director: director}
						proxy.ServeHTTP(c.Writer, c.Request)
					}
					authBool = true
				}
				continue
			}
			if authBool == false {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "白名单鉴权失败",
					"statusCode": "M100001",
					"message1":   "白名单鉴权失败",
					"message2":   wpPackage + " 未在白名单内！",
					"message3":   "请到MengCDN后台添加此仓库信息，或联系管理员处理。",
					"userIp":     UserIP,
				})
			}
		} else {

			// 状态码检测
			if tools.CheckHttpCode1(target+wpProxyPath) == "SM00002" {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "未寻找到文件",
					"statusCode": "404",
					"message1":   "未寻找到文件",
					"message2":   "在这个NPM仓库中未寻找到这个文件。",
					"message3":   "请检查你的路径信息。",
					"userIp":     UserIP,
				})
			} else if tools.CheckHttpCode1(target+wpProxyPath) == "SM00001" {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "代理模块异常",
					"statusCode": "M100002",
					"message1":   "代理模块异常",
					"message2":   "服务器与互联网连接异常！",
					"message3":   "联系管理员处理。",
					"userIp":     UserIP,
				})
			} else {
				// 反向代理模块
				director := func(req *http.Request) {
					req.URL.Scheme = "https"
					req.URL.Host = target
					req.URL.Path = wpProxyPath
					req.Host = target
				}
				proxy := httputil.ReverseProxy{Director: director}
				proxy.ServeHTTP(c.Writer, c.Request)
			}

		}
	}
}

func WpProxyPl(c *gin.Context) {
	if GetCDNSWC("WPpl") == "on" {
		UserIP := c.ClientIP()
		wpPackage := c.Param("package")
		wpPath := c.Param("path")
		//eg1:https://plugins.svn.wordpress.org/365projectorg-widget/
		target := "plugins.svn.wordpress.org"
		wpProxyPath := "/" + wpPackage + "/" + wpPath

		var whiteListSW = GetWLSWC("WPpl")

		if whiteListSW == "on" {
			// 白名单检测
			var authBool bool = false
			for _, i := range GetWLC("WPplWL") {
				if i == wpPackage {
					// 状态码检测
					if tools.CheckHttpCode1(target+wpProxyPath) == "SM00002" {
						c.HTML(http.StatusOK, "error.html", gin.H{
							"title":      "未寻找到文件",
							"statusCode": "404",
							"message1":   "未寻找到文件",
							"message2":   "在这个WP仓库中未寻找到这个文件。",
							"message3":   "请检查你的路径信息。",
							"userIp":     UserIP,
						})
					} else if tools.CheckHttpCode1(target+wpProxyPath) == "SM00001" {
						c.HTML(http.StatusOK, "error.html", gin.H{
							"title":      "代理模块异常",
							"statusCode": "M100002",
							"message1":   "代理模块异常",
							"message2":   "服务器与互联网连接异常！",
							"message3":   "联系管理员处理。",
							"userIp":     UserIP,
						})
					} else {
						// 反向代理模块
						director := func(req *http.Request) {
							req.URL.Scheme = "https"
							req.URL.Host = target
							req.URL.Path = wpProxyPath
							req.Host = target
						}
						proxy := httputil.ReverseProxy{Director: director}
						proxy.ServeHTTP(c.Writer, c.Request)
					}
					authBool = true
				}
				continue
			}
			if authBool == false {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "白名单鉴权失败",
					"statusCode": "M100001",
					"message1":   "白名单鉴权失败",
					"message2":   wpPackage + " 未在白名单内！",
					"message3":   "请到MengCDN后台添加此仓库信息，或联系管理员处理。",
					"userIp":     UserIP,
				})
			}
		} else {

			// 状态码检测
			if tools.CheckHttpCode1(target+wpProxyPath) == "SM00002" {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "未寻找到文件",
					"statusCode": "404",
					"message1":   "未寻找到文件",
					"message2":   "在这个WP仓库中未寻找到这个文件。",
					"message3":   "请检查你的路径信息。",
					"userIp":     UserIP,
				})
			} else if tools.CheckHttpCode1(target+wpProxyPath) == "SM00001" {
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "代理模块异常",
					"statusCode": "M100002",
					"message1":   "代理模块异常",
					"message2":   "服务器与互联网连接异常！",
					"message3":   "联系管理员处理。",
					"userIp":     UserIP,
				})
			} else {
				// 反向代理模块
				director := func(req *http.Request) {
					req.URL.Scheme = "https"
					req.URL.Host = target
					req.URL.Path = wpProxyPath
					req.Host = target
				}
				proxy := httputil.ReverseProxy{Director: director}
				proxy.ServeHTTP(c.Writer, c.Request)
			}

		}
	}
}
