package main

import (
	"net/http"
	"web"
)

func main() {
	// init engine
	engine := web.New()

	// 输入 html
	engine.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Web")
	})

	engine.GET("/hello", func(c *web.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	engine.POST("/login", func(c *web.Context) {
		c.JSON(http.StatusOK, web.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":9999")
}
