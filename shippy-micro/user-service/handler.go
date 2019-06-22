// @Time    : 2019-06-14 19:30
// @Author  : victor！！
// @FileName: handler.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"golang.org/x/crypto/bcrypt"
	userPKG "shippy-micro/user-service/proto/user"
)

type handler struct {
	repo         Repository
	tokenService Authable
	//PubSub       broker.Broker
	Publisher micro.Publisher
}

func (this *handler) Create(ctx context.Context, req *userPKG.User, rsp *userPKG.Response) error {
	// 哈希处理用户密码
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPwd)
	if err := this.repo.Create(req); err != nil {
		return err
	}
	rsp.Created = true

	// 发布带有用户信息的消息
	if err := this.Publisher.Publish(ctx, req); err != nil {
		return err
	}

	return nil
}

// 发布消息通知
//func (this *handler) publishEvent(user *userPKG.User) error {
//	body, err := json.Marshal(user)
//	if err != nil {
//		return err
//	}
//	msg := &broker.Message{
//		Header: map[string]string{
//			"id": user.Id,
//		},
//		Body: body,
//	}
//	log.Printf("====msg:%s\n", msg)
//	// 发布 user.created topic 消息
//	err = this.PubSub.Publish(topic, msg)
//	return err
//}

func (this *handler) Get(ctx context.Context, req *userPKG.User, rsp *userPKG.Response) error {
	user, err := this.repo.Get(req.Id)
	if err != nil {
		return err
	}
	rsp.User = user
	return nil
}

func (this *handler) GetAll(ctx context.Context, req *userPKG.Request, rsp *userPKG.Response) error {
	users, err := this.repo.GetAll()
	if err != nil {
		return err
	}
	rsp.Users = users
	return nil
}

func (this *handler) Auth(ctx context.Context, req *userPKG.User, rsp *userPKG.Token) error {
	// 在part3中直接传参数 &userPKG.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// req.Password 和 user.Password 都是加密后的密码
	// 将无法通过验证
	user, err := this.repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	// 密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}
	token, err := this.tokenService.Encode(user)
	if err != nil {
		return err
	}
	rsp.Token = token
	return nil
}

func (this *handler) ValidateToken(ctx context.Context, req *userPKG.Token, rsp *userPKG.Token) error {
	// Decode token
	claims, err := this.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	rsp.Valid = true
	return nil
}
