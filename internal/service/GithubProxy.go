package service

import (
	"MengCDN/internal/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"strings"
)

func GithubProxy(c *gin.Context) {
	if GetCDNSWC("Github") == "on" {
		UserIP := c.ClientIP()
		owner := c.Param("owner")
		rb := c.Param("repo")
		repo := strings.Split(rb, "@")[0]
		branch := strings.Split(rb, "@")[1]
		path := strings.TrimRight(c.Param("path"), "/")
		//path := c.Param("path")
		githubProxyPath := "/" + owner + "/" + repo + "/" + branch + path
		target := "raw.githubusercontent.com"

		var whiteListSW = GetWLSWC("Github")

		if whiteListSW == "on" {
			// 白名单检测
			var authBool bool = false
			for _, i := range GetWLC("GithubWL") {
				if i == owner+"/"+repo {
					// 状态码检测
					if tools.CheckHttpCode(target+githubProxyPath) == "SM00002" {
						c.Redirect(http.StatusFound, "/browseGH/gh/"+owner+"/"+repo+"@"+branch+"?ghPath="+path)
					} else if tools.CheckHttpCode(target+githubProxyPath) == "SM00001" {
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
							req.URL.Path = githubProxyPath
							req.Host = target
							req.Header.Del("Access-Control-Allow-Origin")
							req.Header.Del("Access-Control-Allow-Headers")
						}
						proxy := httputil.ReverseProxy{Director: director}
						//修改响应处理
						proxy.ModifyResponse = func(response *http.Response) error {
							// 检查原始响应的Content-Type是否为"text/plain"
							if response.Header.Get("Content-Type") == "text/plain; charset=utf-8" {
								// 修改Content-Type为"application/octet-stream"，这是通常不会受到CORB限制的类型
								//response.Header.Set("Content-Type", "application/octet-stream; charset=utf-8")
								lastPath := strings.LastIndex(path, ".")
								if lastPath >= 0 {
									//lastPath1 := path[:lastPath]
									lastPath2 := path[lastPath+len("."):]
									fmt.Print(lastPath2)
									if lastPath2 == "js" {
										response.Header.Set("Content-Type", "application/javascript; charset=utf-8")
									} else if lastPath2 == "css" {
										response.Header.Set("Content-Type", "text/css; charset=utf-8")
									} else {
										response.Header.Set("Content-Type", "application/octet-stream")
										response.Header.Del("Access-Control-Allow-Origin")
									}
								} else {
									response.Header.Set("Content-Type", "application/octet-stream")
									response.Header.Del("Access-Control-Allow-Origin")
								}
							}
							return nil
						}
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
					"message2":   owner + "/" + repo + " 未在白名单内！",
					"message3":   "请到MengCDN后台添加此仓库信息，或联系管理员处理。",
					"userIp":     UserIP,
				})
			}
		} else {
			// 状态码检测
			if tools.CheckHttpCode(target+githubProxyPath) == "SM00002" {
				c.Redirect(http.StatusFound, "/browseGH/gh/"+owner+"/"+repo+"@"+branch+"?ghPath="+path)
			} else if tools.CheckHttpCode(target+githubProxyPath) == "SM00001" {
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
					req.URL.Path = githubProxyPath
					req.Host = target
					req.Header.Del("Access-Control-Allow-Origin")
					req.Header.Del("Access-Control-Allow-Headers")
				}
				proxy := httputil.ReverseProxy{Director: director}
				//修改响应处理
				proxy.ModifyResponse = func(response *http.Response) error {
					// 检查原始响应的Content-Type是否为"text/plain"
					if response.Header.Get("Content-Type") == "text/plain; charset=utf-8" {
						// 修改Content-Type为"application/octet-stream"，这是通常不会受到CORB限制的类型
						//response.Header.Set("Content-Type", "application/octet-stream; charset=utf-8")
						lastPath := strings.LastIndex(path, ".")
						if lastPath >= 0 {
							//lastPath1 := path[:lastPath]
							lastPath2 := path[lastPath+len("."):]
							fmt.Print(lastPath2)
							if lastPath2 == "js" {
								response.Header.Set("Content-Type", "application/javascript; charset=utf-8")
							} else if lastPath2 == "css" {
								response.Header.Set("Content-Type", "text/css; charset=utf-8")
							} else {
								response.Header.Set("Content-Type", "application/octet-stream")
								response.Header.Del("Access-Control-Allow-Origin")
							}
						} else {
							response.Header.Set("Content-Type", "application/octet-stream")
							response.Header.Del("Access-Control-Allow-Origin")
						}
					}
					return nil
				}
				proxy.ServeHTTP(c.Writer, c.Request)
			}
		}
	}
}
