// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: func.go
// @Date: 2022/5/13 14:12
// @Desc:
package shares

import (
	"context"
	"github.com/stonebirdjx/stonebird-server/conf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createUniqueIndex(collection string) int {
	mc, err := NewMongoClient()
	if err != nil {
		return 1
	}

	col := mc.Database(conf.Database).Collection(collection)
	indexName := "index_by_ip"
	unique := true
	_, err = col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"ip", 1}},
		Options: &options.IndexOptions{
			Name:   &indexName,
			Unique: &unique,
		},
	})

	if err != nil {
		return 1
	}
	return 0
}
