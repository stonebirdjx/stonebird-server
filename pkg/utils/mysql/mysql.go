// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: mysql.go  2022/6/19 11:01.
// Desc:

// Package mysql connection util.
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stonebirdjx/stonebird-server/configs"
)

var conn *sql.DB

// Dial conn mysql server.
func Dial() (*sql.DB, error) {
	var err error
	if conn != nil {
		return conn, err
	}

	conn, err = sql.Open(configs.MysqlDriver, configs.MysqlDsn())
	return conn, err
}
