// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: main.go  2022/6/18 10:55.
// Desc:

package main

import (
	"github.com/stonebirdjx/stonebird-server/pkg/apis"
	"log"
)

func main() {
	err := apis.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
