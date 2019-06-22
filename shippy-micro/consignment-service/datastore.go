// @Time    : 2019-06-14 10:52
// @Author  : victor！！
// @FileName: datastore.go
// @Software: IntelliJ IDEA

package main

import (
	"gopkg.in/mgo.v2"
)

func CreateSession(host string) (*mgo.Session, error) {
	sess, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	sess.SetMode(mgo.Monotonic, true)
	sess.SetPoolLimit(300)
	return sess, nil
}
