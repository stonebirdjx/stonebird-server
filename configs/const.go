// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: const.go  2022/6/18 11:39.
// Desc:

package configs

import (
	"fmt"
	"os"
)

// system info
const (
	ServerName = "stone bird server v2.0.1/gin"
	Email      = "<124586360@qq.com,g1245863260@gmail.com>"
	APIv1      = "/api/v1"
)

// database info
const (
	MysqlDriver = "mysql"
	mysqlHost   = "STB_MYSQL_HOST"
	mysqlPort   = "STB_MYSQL_PORT"
	mysqlUser   = "STB_MYSQL_USER"
	mysqlPasswd = "STB_MYSQL_PASSWD"
	mysqlName   = "stonebird"
)

// MysqlDsn dbUser:dbPasswd@tcp(dbHost:dbPort)/dbName.
func MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv(mysqlUser),
		os.Getenv(mysqlPasswd),
		os.Getenv(mysqlHost),
		os.Getenv(mysqlPort),
		os.Getenv(mysqlName))
}
