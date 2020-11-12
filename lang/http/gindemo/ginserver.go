package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
1. gin配置
* gin-gonic/gin：https://github.com/gin-gonic/gin
* go get -u github.com/gin-gonic/gin

2. gin engine用法
* engine:=gin.Default()：创建gin引擎
* engine.Use()：gin中间件
* engine.Get()

3. gin context用法
* context：gin.Context是一个实现go标准库Context接口的结构体
* middleware：注册中间件函数，使得所有请求均经过这些middleware，使得AOP的用法简化，省去annotation注解的写法
* c.JSON()
* c.String()

2.其他第三方库的协作
* zap：go get -u go.uber.org/zap
*/

const keyRequest = "requestId"

func main() {
	r := gin.Default()
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	//middleware用法
	r.Use(func(c *gin.Context) {
		s := time.Now()
		c.Next()

		//path,log latency, response code
		log.Info("Incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(c *gin.Context) {
		//为每个进入的request添加一个requestId
		//其他的context及http请求均可访问该requestId
		c.Set(keyRequest, rand.Int())
		c.Next()
	})

	//http GET用法
	r.GET("/ping", func(c *gin.Context) {

		//http response
		h := gin.H{
			"message": "pong",
		}
		//http response，设置request的key-value pair
		if rid, exists := c.Get(keyRequest); exists {
			h[keyRequest] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
