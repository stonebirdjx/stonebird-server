// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: midderwares.go
// @Date: 2022/5/12 13:57
// @Desc:
package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/conf"
	"net/http"
)

// Headers set the response header middleware.
func Headers(c *gin.Context) {
	c.Header("Serve", conf.Serve)
	c.Header("Email", conf.Email)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "false")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusNoContent, nil)
	}
	c.Next()
}
