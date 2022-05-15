// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: users.go
// @Date: 2022/5/12 16:55
// @Desc:
package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	name := c.Param("name")
	fmt.Println("name", name)
	c.JSON(http.StatusOK, gin.H{
		"user_name": "Stone Bird",
		"nick_name": "石鸟路遇",
		"position":  "中国-广东省-深圳市",
	})
}
