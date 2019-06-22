// @Time    : 2019-06-14 13:38
// @Author  : victor！！
// @FileName: datastore.go
// @Software: IntelliJ IDEA

package main

import (
	"errors"
	"gopkg.in/mgo.v2"
)

func CreateSession(host string) (*mgo.Session, error) {
	if host == "" {
		return nil, errors.New("dial mongo host but host is empty")
	}
	sess, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	sess.SetMode(mgo.Monotonic, true)
	sess.SetPoolLimit(300)
	return sess, nil
}
