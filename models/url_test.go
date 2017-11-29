package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func init() {
	db_url := "127.0.0.1:3000"
	session, err := mgo.Dial(db_url)
	if err != nil {
		panic(err)
	}
	_db = session.DB("short_url_test")
}

func TestUrl(t *testing.T) {
	u := Url{
		Id:        1,
		SourceUrl: "www.qq.com",
	}

	err := u.Insert()
	assert.Nil(t, err)

	u2 := Url{
		Id:        2,
		SourceUrl: "www.bilibili.com",
	}
	u2.Insert()

	u.SourceUrl = "www.baidu.com"
	u.Update()

	uu := Url{}
	uu.Id = 2
	uu.FindById()

	assert.Equal(t, "www.bilibili.com", uu.SourceUrl)

	uu.DeleteById()
}

func TestGenId(t *testing.T) {
	url := &Url{}
	url.SourceUrl = "http://www.facebook3.com"
	err := url.GenId()
	assert.Nil(t, err)

	err = url.Save()
	assert.Nil(t, err)
}
