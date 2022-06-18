// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: story.go  2022/6/18 11:44.
// Desc:

package story

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "a little story",
	})
}
