// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: router.go
// @Date: 2022/5/12 14:02
// @Desc:
package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/biz/api/v1/stories"
	"github.com/stonebirdjx/stonebird-server/biz/api/v1/themes"
	"github.com/stonebirdjx/stonebird-server/biz/api/v1/users"
	"github.com/stonebirdjx/stonebird-server/middlewares"
	"net/http"
)

func engine() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Headers)
	// gin 注册 pprof
	pprof.Register(r)
	consumer(r)
	return r
}

func consumer(r *gin.Engine) {
	r.Any("/", root)
	r.GET("/ping", Ping)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("stories", stories.List)
		apiV1.POST("stories", stories.Create)
		apiV1.GET("stories/:name", stories.Get)
		apiV1.DELETE("stories/:name", stories.Delete)
		apiV1.POST("stories/email", stories.Email)
		apiV1.GET("users/:name", users.Get)
		apiV1.POST("themes", themes.Create)
		apiV1.GET("themes", themes.Get)
		apiV1.PUT("themes", themes.Put)
	}
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to visit stone bird Backend ,but please visit https://www.hjxstbserver.xyz through your browser",
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Engine() *gin.Engine {
	return engine()
}
