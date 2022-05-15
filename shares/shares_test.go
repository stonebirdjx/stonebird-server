// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: shares_test.go
// @Date: 2022/5/13 2:48
// @Desc:
package shares

import (
	"github.com/stonebirdjx/stonebird-server/conf"
	"testing"
)

func TestCreateUniqueIndex(t *testing.T) {
	createUniqueIndex(conf.ThemesCollection)
}
