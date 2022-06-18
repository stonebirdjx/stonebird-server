// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: apis.go  2022/6/18 11:02.
// Desc:

// Package apis Define the api route in the backend.
package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/configs"
	v1 "github.com/stonebirdjx/stonebird-server/pkg/apis/v1"
	"github.com/stonebirdjx/stonebird-server/pkg/handler"
)

func engine() *gin.Engine {
	r := gin.Default()

	// router detail
	r.GET("/ping", handler.Ping)
	v1.APIv1Consumer(r)
	return r
}

func Serve() error {
	r := engine()
	return r.Run(configs.GetPort())
}
