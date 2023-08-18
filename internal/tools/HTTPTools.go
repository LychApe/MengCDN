package tools

import (
	"io"
	"net/http"
)

// CheckHttpCode 状态码检测组件
func CheckHttpCode(path string) string {
	resp, err := http.Get("https://" + path)

	if err != nil {
		println("==========================================================")
		println("[代理模块自检异常]:[GithubProxy][服务器与互联网连接异常！]")
		println("==========================================================")
		return "SM00001"
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				println(err)
			}
		}(resp.Body)

		if resp.StatusCode != 200 {
			return "SM00002"
		} else {
			return "SM00003"
		}
	}

}

// CheckHttpCode1 状态码检测组件
func CheckHttpCode1(path string) string {
	resp, err := http.Get("https://" + path)

	if err != nil {
		println("==========================================================")
		println("[代理模块自检异常]:[GithubProxy][服务器与互联网连接异常！]")
		println("==========================================================")
		return "SM00001"
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				println(err)
			}
		}(resp.Body)

		if resp.StatusCode != 200 {
			return "SM00002"
		} else {
			return "SM00003"
		}
	}
}

// CheckHttpCode2 状态码检测组件
func CheckHttpCode2(path string) string {
	resp, err := http.Get("https://" + path)

	if err != nil {
		println("==========================================================")
		println("[代理模块自检异常]:[GithubProxy][服务器与互联网连接异常！]")
		println("==========================================================")
		return "SM00001"
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				println(err)
			}
		}(resp.Body)

		if resp.StatusCode != 200 {
			return "SM00002"
		} else if resp.StatusCode == 302 {
			return "SM00004"
		} else {
			return "SM00003"
		}
	}
}
