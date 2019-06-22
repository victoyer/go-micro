// @Time    : 2019-06-14 13:41
// @Author  : victor！！
// @FileName: repository.go
// @Software: IntelliJ IDEA

package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	vessPKG "shippy-micro/vessel-service/proto/vessel"
)

type IRepository interface {
	FindAvailable(*vessPKG.Specification) (*vessPKG.Vessel, error)
	Create(*vessPKG.Vessel) error
	Close()
}

type Repository struct {
	sess *mgo.Session
}

func (this *Repository) collection() *mgo.Collection {
	return this.sess.DB(DB_NAME).C(CON_COLLECTION)
}

func (this *Repository) FindAvailable(spec *vessPKG.Specification) (*vessPKG.Vessel, error) {
	var vessel *vessPKG.Vessel
	err := this.collection().Find(
		bson.M{
			"capacity":  bson.M{"$gte": spec.Capacity},
			"maxweight": bson.M{"$gte": spec.MaxWeight},
		}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (this *Repository) Create(vessel *vessPKG.Vessel) error {
	return this.collection().Insert(vessel)
}

func (this *Repository) Close() {
	this.sess.Close()
}
