// Copyright (c) 2018 hu. All rights reserved.
// Personal code, only for learning and communication.
// You can contact me in the following ways:
//    stonebirdjx <1245863260@qq.com, g1245863260@gmail.com>
//    https://www.hjxstbserver.xyz/
// File: router.go  2022/6/18 11:35.
// Desc:

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/configs"
	"github.com/stonebirdjx/stonebird-server/pkg/handler/story"
)

// APIv1Consumer v1 router detail.
func APIv1Consumer(r *gin.Engine) {
	v1Group := r.Group(configs.APIv1)
	{
		storyAPI(v1Group)
	}
}

func storyAPI(rg *gin.RouterGroup) {
	rg.GET("/story", story.Get)
}
