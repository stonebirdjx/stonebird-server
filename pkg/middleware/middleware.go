// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: middleware.go  2022/6/18 11:59.
// Desc:

// Package middleware for server app.
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/configs"
)

func DeclareHeader(c *gin.Context) {
	c.Header("Serve", configs.ServerName)
	c.Header("Email", configs.Email)
	c.Next()
}
