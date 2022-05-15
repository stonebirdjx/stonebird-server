// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: main.go
// @Date: 2022/5/12 13:43
// @Desc:
package main

import (
	"context"
	"github.com/stonebirdjx/stonebird-server/conf"
	"github.com/stonebirdjx/stonebird-server/routers"
	"github.com/stonebirdjx/stonebird-server/shares"
	"log"
)

func main() {
	defer func() {
		if mc, _ := shares.NewMongoClient(); mc != nil {
			_ = mc.Disconnect(context.Background())
		}
	}()

	r := routers.Engine()
	if err := r.Run(conf.Port); err != nil {
		log.Fatal(err)
	}
}
