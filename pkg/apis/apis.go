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
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/configs"
	v1 "github.com/stonebirdjx/stonebird-server/pkg/apis/v1"
	"github.com/stonebirdjx/stonebird-server/pkg/handler"
	"github.com/stonebirdjx/stonebird-server/pkg/middleware"
	"github.com/stonebirdjx/stonebird-server/pkg/utils/mysql"
)

func engine() *gin.Engine {
	r := gin.Default()
	// gin use pprof
	pprof.Register(r)

	r.Use(middleware.DeclareHeader)
	// router detail
	r.GET("/ping", handler.Ping)
	v1.APIv1Consumer(r)
	return r
}

func Serve() error {
	if err := mysqlExamine(); err != nil {
		return err
	}

	r := engine()
	return r.Run(configs.GetPort())
}

// mysqlExamine test mysql conn and set maxlifetime and idleconns.
func mysqlExamine() error {
	conn, err := mysql.Dial()
	if err != nil {
		return err
	}
	if err = conn.Ping(); err != nil {
		return err
	}
	conn.SetConnMaxLifetime(10)
	conn.SetMaxIdleConns(5)
	return err
}
