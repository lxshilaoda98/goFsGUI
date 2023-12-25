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
			"title": "Main website",
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
