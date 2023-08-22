package router

import (
	v1 "MengCDN/internal/api/v1"
	"MengCDN/internal/middleware"
	"MengCDN/internal/service"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("indexF", "web/front/Index/dist/index.html")
	r.AddFromFiles("indexF1", "web/front/Index/dist/start.html")
	r.AddFromFiles("indexF2", "web/front/Index/dist/introduction.html")
	r.AddFromFiles("indexF3", "web/front/Index/dist/404.html")
	r.AddFromFiles("admin", "web/front/MengCDN_Admin/dist/index.html")
	r.AddFromFiles("gh", "web/front/MengCDN_Front/dist/index.html")
	return r
}

func InitRouter() {
	router := gin.Default()

	router.HTMLRender = createMyRender()
	router.Use(middleware.CORSMiddleware())

	//[WebUI][Error]
	router.Static("MengCDNAdmin/assets", "web/front/MengCDN_Admin/dist/assets")
	router.Static("assets", "web/front/Index/dist/assets")
	router.Static("browseGH/assets", "web/front/MengCDN_Front/dist/assets")
	router.LoadHTMLGlob("web/error/error.html")

	AuthRouterV1 := router.Group("api/v1")
	AuthRouterV1.Use(middleware.JwtToken())
	{
		// 编辑白名单
		AuthRouterV1.PUT("/cdnWL/:mk", v1.PUTCdnWL)
		// 编辑CDN模块开关
		AuthRouterV1.PUT("/cdnSW/:id/:mk/:sw", v1.PUTCdnSW)
	}
	PublicRouterV1 := router.Group("api/v1")
	{
		//查询白名单
		PublicRouterV1.POST("/cdnWL/:mk", v1.CdnWL)
		//查询CDN模块开关
		PublicRouterV1.POST("/cdnSW/:id/:mk", v1.CdnSW)
		// 登录
		PublicRouterV1.POST("login", v1.Login)
	}

	//[v1][Github]
	githubProxyRouter := router.Group("/gh")
	{
		githubProxyRouter.GET("/:owner/:repo/*path", service.GithubProxy)
	}
	//[GithubList UI]
	router.GET("/browseGH", func(c *gin.Context) {
		c.HTML(200, "gh", nil)
		content, err := ioutil.ReadFile("web/front/MengCDN_Front/dist/index.html")
		if (err) != nil {
			UserIP := c.ClientIP()
			c.HTML(http.StatusOK, "error.html", gin.H{
				"title":      "查询错误",
				"statusCode": "M100003",
				"message1":   "查询错误",
				"message2":   "查询路径信息错误",
				"message3":   "请检查你的路径信息。正确方法: /用户名/仓库@分支?ghPath=/要查询的仓库内目录",
				"userIp":     UserIP,
			})
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
		router.NoRoute(func(c *gin.Context) {
			content, err := ioutil.ReadFile("web/front/MengCDN_Front/dist/index.html")
			if (err) != nil {
				UserIP := c.ClientIP()
				c.HTML(http.StatusOK, "error.html", gin.H{
					"title":      "查询错误",
					"statusCode": "M100003",
					"message1":   "查询错误",
					"message2":   "查询路径信息错误",
					"message3":   "请检查你的路径信息。正确方法: /用户名/仓库/分支?ghPath=/要查询的仓库内目录",
					"userIp":     UserIP,
				})
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		})
	})

	//[v1][NPM]
	npmProxyRouter := router.Group("/npm")
	{
		npmProxyRouter.GET("/:package/:version/*path", service.NpmProxy)
		npmProxyRouter.GET("/browse/:package/:version/*path", service.NpmProxyBrowse)
	}
	router.GET("/browse/:ab/*path", func(c *gin.Context) {
		ab := c.Param("ab")
		ab1 := strings.Split(ab, "@")[0]
		ab2 := strings.Split(ab, "@")[1]
		npmPath := c.Param("package")
		c.Redirect(http.StatusFound, "/npm/browse/"+ab1+"/"+ab2+"/"+npmPath)
	})

	//[v1][WordPress]
	wpProxyRouter := router.Group("/wp")
	{
		wpProxyRouter.GET("/theme/:package/:version/*path", service.WpProxyTh)
		wpProxyRouter.GET("/plugin/:package/*path", service.WpProxyPl)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "indexF", nil)
		content, err := ioutil.ReadFile("web/front/Index/dist/index.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
	})
	router.GET("/start.html", func(c *gin.Context) {
		c.HTML(200, "indexF1", nil)
		content, err := ioutil.ReadFile("web/front/Index/dist/start.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
	})
	router.GET("/introduction.html", func(c *gin.Context) {
		c.HTML(200, "indexF2", nil)
		content, err := ioutil.ReadFile("web/front/Index/dist/introduction.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
	})
	router.GET("/404.html", func(c *gin.Context) {
		c.HTML(200, "indexF3", nil)
		content, err := ioutil.ReadFile("web/front/Index/dist/404.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
	})

	router.GET("/MengCDNAdmin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
		content, err := ioutil.ReadFile("web/front/MengCDN_Admin/dist/index.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write(content)
		c.Writer.Flush()
		router.NoRoute(func(c *gin.Context) {
			accept := c.Request.Header.Get("Accept")
			flag := strings.Contains(accept, "text/html")
			if flag {
				content, err := ioutil.ReadFile("web/front/MengCDN_Admin/dist/index.html")
				if (err) != nil {
					c.Writer.WriteHeader(404)
					c.Writer.WriteString("Not Found")
					return
				}
				c.Writer.WriteHeader(200)
				c.Writer.Header().Add("Accept", "text/html")
				c.Writer.Write(content)
				c.Writer.Flush()
			}
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(301, "/browseGH/")
	})

	err := router.Run(":8001")
	if err != nil {
		return
	}

}
