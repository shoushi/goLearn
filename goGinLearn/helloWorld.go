package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 方法一：直接写map
func sayHello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

// 方法二：写struct,结构体 首字母大写，保证JSON可以取到
type message struct {
	//可以使用tag对结构体字段进行自定义操作
	Name string `json:"name"`
	Age  int
}

func sayHello2(context *gin.Context) {
	// 获取url 参数
	name := context.DefaultQuery("name", "guanliqun")
	age, _ := strconv.Atoi(context.Query("age"))
	context.JSON(http.StatusOK, message{
		Name: name,
		Age:  age,
	})
}

// 支持中间件执行在请求前
func logSave(c *gin.Context) {
	fmt.Println("get request")
}
func main() {
	// 创建默认路由
	r := gin.Default()
	// 全局中间件
	// r.Use(logSave)
	r.GET("/hello", logSave, sayHello2)
	r.GET("/baidu", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "https://www.rushb.io/")
	})
	group := r.Group("group")
	group.Use(logSave)
	{
		group.GET("sogo", func(context *gin.Context) {
			//路由重定向
			context.Request.URL.Path = "/baidu"
			r.HandleContext(context)
		})
		group.Any("any", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"test": 121,
			})
		})
	}

	mock(r)
	r.Run(":8080")
}
