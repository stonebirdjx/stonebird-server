// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: mysql_test.go  2022/6/19 11:31.
// Desc:

package mysql

import "testing"

func TestDial(t *testing.T) {
	conn, err := Dial()
	if err != nil {
		t.Fatal("open mysql dsn failed")
	}
	if err := conn.Ping(); err != nil {
		t.Fatal("ping mysql failed")
	}
	t.Log("conn mysql driver success")
}
