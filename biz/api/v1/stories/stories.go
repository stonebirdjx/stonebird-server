// Copyright (c) 2021 hu. All rights reserved.
// @Author: stonebirdjx
// @Email: 1245863260@qq.com, g1245863260@gmail.com
// @File: stories.go
// @Date: 2022/5/12 15:06
// @Desc:
package stories

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stonebirdjx/stonebird-server/conf"
	"github.com/stonebirdjx/stonebird-server/shares"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type pageInfo struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Avatar   string             `json:"avatar" bson:"avatar"`       // 图标文字
	Name     string             `json:"name" bson:"name"`           // 用户名
	NickName string             `json:"nick_name" bson:"nick_name"` // 昵称
	Position string             `json:"position" bson:"position"`   // 位置
	Email    string             `json:"email" bson:"email"`         // 邮箱
	Gender   string             `json:"gender" bson:"gender"`       // 性别
	Hobby    string             `json:"hobby" bson:"hobby"`         // 爱好
	Title    []string           `json:"title" bson:"title"`         // 头衔
	Story    string             `json:"story" bson:"story"`         // 故事
}

func List(c *gin.Context) {
	col, err := shares.ConnCol(conf.Database, conf.StoryCollection, c)
	if err != nil {
		return
	}

	var pis []pageInfo
	cursor, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to query story list",
			"err":     err.Error(),
		})
		return
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var pi pageInfo
		err := cursor.Decode(&pi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Decoding story list failed",
				"err":     err.Error(),
			})
			return
		}
		pis = append(pis, pi)
	}

	c.JSON(http.StatusOK, gin.H{
		"total":    len(pis),
		"metadata": pis,
	})
}

func Get(c *gin.Context) {
	name := c.Param("name")
	name, _ = url.QueryUnescape(name)
	col, err := shares.ConnCol(conf.Database, conf.StoryCollection, c)
	if err != nil {
		return
	}

	var pi pageInfo

	if err := col.FindOne(context.Background(), bson.D{{"name", name}}).Decode(&pi); err != nil {
		c.JSON(http.StatusOK, pageInfo{
			Id:       primitive.NewObjectID(),
			Avatar:   "TU",
			Name:     "Test User",
			NickName: "虚拟用户",
			Position: "中国-广东省-深圳市",
			Gender:   "男",
			Hobby:    "唱歌、听音乐、看电影、看韩剧、看综艺娱乐节目、看书、看小说、看杂志、逛街、购物",
			Title:    nil,
			Story:    "故事书写中",
		})
		return
	}

	c.JSON(http.StatusOK, pi)
}

func Create(c *gin.Context) {
	var pi pageInfo
	if err := c.BindJSON(&pi); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The data conversion failed, the content content cannot be obtained",
			"err":     err.Error(),
		})
		return
	}

	col, err := shares.ConnCol(conf.Database, conf.StoryCollection, c)
	if err != nil {
		return
	}

	if strings.TrimSpace(pi.Name) == conf.EmptyString {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Name cant not be  blank",
			"err":     "Name is nil",
		})
		return
	}

	pi.Id = primitive.NewObjectID()
	res, err := col.InsertOne(context.TODO(), pi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Create user success %s:%s", pi.Name, res.InsertedID),
	})
}

func Delete(c *gin.Context) {
	name := c.Param("name")
	name, _ = url.QueryUnescape(name)
	col, err := shares.ConnCol(conf.Database, conf.StoryCollection, c)
	if err != nil {
		return
	}

	_, err = col.DeleteOne(context.Background(), bson.D{{"name", name}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete story",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func Email(c *gin.Context) {
	clientIp := c.ClientIP()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Get Msg content err",
			"err":     err.Error(),
		})
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", conf.EmailSend)
	m.SetHeader("To", conf.EmailRecv)
	m.SetHeader("Subject", fmt.Sprintf("Story页面消息源自%s", clientIp))
	m.SetBody("text/html", string(b))

	gomailDialer := gomail.NewDialer(conf.EmailHost,
		conf.EmailPort,
		conf.EmailSend,
		conf.EmailAuthorization)

	err = gomailDialer.DialAndSend(m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to send email",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Send msg success",
	})

}
