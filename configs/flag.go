// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: flag.go  2022/6/18 11:17.
// Desc:

package configs

import (
	"flag"
	"fmt"
	"os"
)

var (
	port    = flag.String("port", "6610", "enter the server port")
	h       = flag.Bool("h", false, "print help text")
	help    = flag.Bool("help", false, "print help text")
	v       = flag.Bool("v", false, "print server version")
	version = flag.Bool("version", false, "print server version")
)

const serverVersion = "stone bird server v2.0.1"

func init() {
	flag.Parse()

	// show help
	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	// show version
	if *v || *version {
		fmt.Println(serverVersion)
		os.Exit(0)
	}

}

func GetPort() string {
	return ":" + *port
}
