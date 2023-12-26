package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE ,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func decodefsPass(base64Sub string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Sub)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return ""
	}

	// 将字节转换为字符串
	decodedString := string(decodedBytes)

	unescapedString, err := url.QueryUnescape(decodedString)
	if err != nil {
		fmt.Println("Error unescaping string:", err)
		return ""
	}

	return unescapedString

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Use(cors())
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bash.html", gin.H{
			"title":           "Main website",
			"fs_version":      "version:查看版本号",
			"fs_sofia_status": "sofia status:查看sofia信息",
			//http_cache
			"http_clear_cache": "http_clear_cache:清理http缓存",
			//http_xml_curl
			"xml_flush_cache":        "xml_flush_cache:清理xml缓存",
			"xml_flush_cache_sipone": "xml_flush_cache id 1002 domain-name:清理单个缓存",
			//mod_sofia
			"fsctl_flush_db_handles": "fsctl flush_db_handles:关闭不再需要的数据库连接",
			"help":                   "help:帮助",
			//json {"command" : "status", "data" : ""} 返回json格式
			"module_exists": "module_exists 模块名称:检查模块是否已加载",
			"bridgeUser":    "originate user/分机号 &park():呼叫分机，并挂起",
			"bridgetoUser":  "originate user/a分机号 &bridge(user/b分机号):先呼叫a，a接起后呼叫b",
		})
	})

	router.GET("/fscli", func(c *gin.Context) {
		body := c.DefaultQuery("body", "")
		bodyJson := decodefsPass(body)
		if bodyJson != "" {
			var fsGUIInput FsGUIInput
			json.Unmarshal([]byte(bodyJson), &fsGUIInput)
			api, err := FsApi(&fsGUIInput)
			if err != nil {
				c.String(200, err.Error())
				return
			}
			c.String(200, api)
		} else {
			c.String(200, "传递参数异常!")
			return
		}
	})
	router.Run(":1225")
}
