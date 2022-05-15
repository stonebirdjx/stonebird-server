// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: mongo.go
// @Date: 2022/5/13 2:23
// @Desc:
package shares

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

// D系列包含4种类型：
//– D：一个BSON文档。这个类型应该被用在顺序很重要的场景， 比如MongoDB命令。
//– M: 一个无需map。 它和D是一样的， 除了它不保留顺序。
//– A: 一个BSON数组。
//– E: 在D里面的一个单一的子项。
var client *mongo.Client

func newMongoClient() (*mongo.Client, error) {
	var err error
	if client != nil {
		return client, err
	}

	//
	credential := options.Credential{
		Username: conf.MongoUser,
		Password: conf.MongoPassword,
	}
	uri := fmt.Sprintf("%s://%s:%s",
		conf.MongoProtocol,
		conf.RemoteIP,
		conf.MongoPort)
	// 带用户名和密码连接
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err = mongo.Connect(context.TODO(), clientOpts)
	return client, err
}

func NewMongoClient() (*mongo.Client, error) {
	return newMongoClient()
}

func ConnCol(database, collection string, c *gin.Context) (*mongo.Collection, error) {
	mc, err := newMongoClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to connect to mongo db",
			"err":     err.Error(),
		})
		return nil, err
	}

	col := mc.Database(database).Collection(collection)
	return col, err
}
