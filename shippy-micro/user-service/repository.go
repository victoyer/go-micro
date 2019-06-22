// @Time    : 2019-06-14 19:20
// @Author  : victor！！
// @FileName: repository.go
// @Software: IntelliJ IDEA

package main

import (
	"github.com/jinzhu/gorm"
	"log"
	userPKG "shippy-micro/user-service/proto/user"
)

type IRepository interface {
	Create(*userPKG.User) error
	Get(id string) (*userPKG.User, error)
	GetAll() ([]*userPKG.User, error)
	GetByEmail(email string) (*userPKG.User, error)
}

type Repository struct {
	db *gorm.DB
}

func (this *Repository) Create(user *userPKG.User) error {
	if err := this.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (this *Repository) Get(id string) (*userPKG.User, error) {
	var user *userPKG.User
	user.Id = id
	if err := this.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (this *Repository) GetAll() ([]*userPKG.User, error) {
	var users []*userPKG.User
	if err := this.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (this *Repository) GetByEmail(email string) (*userPKG.User, error) {
	user := &userPKG.User{}
	if err := this.db.Where("email = ?", email).Find(user).Error; err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}
