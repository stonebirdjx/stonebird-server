// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: const.go  2022/6/18 11:39.
// Desc:

package configs

import "fmt"

// system info
const (
	ServerName = "stone bird server v2.0.1/gin"
	Email      = "<124586360@qq.com,g1245863260@gmail.com>"
	APIv1      = "/api/v1"
)

// database info
const (
	MysqlDriver = "mysql"
	mysqlHost   = "106.13.178.91"
	mysqlPort   = "3306"
	mysqlUser   = "root"
	mysqlPasswd = "123456"
	mysqlName   = "stonebird"
)

// MysqlDsn dbUser:dbPasswd@tcp(dbHost:dbPort)/dbName.
// 用户名:密码啊@tcp(ip:端口)/数据库的名字
func MysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPasswd, mysqlHost, mysqlPort, mysqlName)
}
