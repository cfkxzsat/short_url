package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Url struct {
	Id        int    `bson:"_id"`
	SourceUrl string `bson:"SourceUrl"`
	ShortUrl  string `bson:"-"`
}

var (
	URL_COLLECTION = "Url"
)

func (url *Url) GenId() error {
	sourceUrl := url.SourceUrl
	err := GetDB().C(URL_COLLECTION).Find(bson.M{"SourceUrl": sourceUrl}).One(url)
	if err != nil {
		url.Id, err = IncrMaxId("url") //??
		if err != nil {
			return err
		}
	}
	return nil
}

func (url *Url) Save() error {
	return url.Upsert()
}

func (url *Url) Insert() error {
	return GetDB().C(URL_COLLECTION).Insert(url)
}

func (url *Url) FindById() error {

	return GetDB().C(URL_COLLECTION).FindId(url.Id).One(url)
}

func (url *Url) Update() error {
	return GetDB().C(URL_COLLECTION).Update(bson.M{"_id": url.Id}, url)
}

func (url *Url) Upsert() error {
	_, err := GetDB().C(URL_COLLECTION).Upsert(bson.M{"_id": url.Id}, url)
	return err
}

func (url *Url) DeleteById() error {
	return GetDB().C(URL_COLLECTION).Remove(bson.M{"_id": url.Id})
}
