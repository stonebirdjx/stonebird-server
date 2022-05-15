// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: themes.go
// @Date: 2022/5/14 12:09
// @Desc:
package themes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/conf"
	"github.com/stonebirdjx/stonebird-server/shares"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type bodyInfo struct {
	Id   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Ip   string             `json:"ip" bson:"ip"`
	Dark bool               `json:"dark" bson:"dark"`
}

func Create(c *gin.Context) {
	var bi bodyInfo
	if err := c.BindJSON(&bi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "themes BindJSON err",
		})
		return
	}

	col, err := shares.ConnCol(conf.Database, conf.ThemesCollection, c)
	if err != nil {
		return
	}

	bi.Id = primitive.NewObjectID()
	ior, err := col.InsertOne(context.Background(), bi)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "themes save mongo err",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "creat success",
		"id":      ior.InsertedID,
	})
}

func Get(c *gin.Context) {
	ip := c.ClientIP()
	col, err := shares.ConnCol(conf.Database, conf.ThemesCollection, c)
	if err != nil {
		return
	}

	var bi bodyInfo
	if err := col.FindOne(context.Background(), bson.D{{"ip", ip}}).Decode(&bi); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"dark": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"dark": bi.Dark,
	})
}

func Put(c *gin.Context) {
	col, err := shares.ConnCol(conf.Database, conf.ThemesCollection, c)
	if err != nil {
		return
	}

	var bi bodyInfo
	if err := c.BindJSON(&bi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "themes BindJSON err",
			"err":     err.Error(),
		})
		return
	}

	clentIp := c.ClientIP()
	var result bson.M
	err = col.FindOne(context.Background(), bson.M{"ip": clentIp}).Decode(&result)
	if len(result) > 0 {
		updateOneFilter := bson.M{"ip": clentIp}
		updateOneSet := bson.M{"$set": bson.M{"dark": bi.Dark}}
		ret, err := col.UpdateOne(context.Background(), updateOneFilter, updateOneSet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "themes update bd err",
				"err":     err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "themes save success",
			"id":      ret.UpsertedID,
		})
		return
	}

	bi.Id = primitive.NewObjectID()
	bi.Ip = clentIp
	res, err := col.InsertOne(context.TODO(), bi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "themes save bd err",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "themes save success",
		"id":      res.InsertedID,
	})

}
