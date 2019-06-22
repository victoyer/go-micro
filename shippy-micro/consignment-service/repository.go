// @Time    : 2019-06-14 11:06
// @Author  : victor！！
// @FileName: repository.go
// @Software: IntelliJ IDEA

package main

import (
	"gopkg.in/mgo.v2"
	consPKG "shippy-micro/consignment-service/proto/consignment"
)

type IRepository interface {
	Create(*consPKG.Consignment) (*consPKG.Consignment, error)
	GetAll() ([]*consPKG.Consignment, error)
	Close()
}

type Repository struct {
	sess *mgo.Session
}

func (this *Repository) collection() *mgo.Collection {
	return this.sess.DB(DB_NAME).C(CON_COLLECTION)
}

func (this *Repository) Create(consignment *consPKG.Consignment) (*consPKG.Consignment, error) {
	if err := this.collection().Insert(consignment); err != nil {
		return nil, err
	}
	return consignment, nil
}

func (this *Repository) GetAll() ([]*consPKG.Consignment, error) {
	var cons []*consPKG.Consignment
	if err := this.collection().Find(nil).All(&cons); err != nil {
		return nil, err
	}
	return cons, nil
}

func (this *Repository) Close() {
	this.sess.Close()
}
