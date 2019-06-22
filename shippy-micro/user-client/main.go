// @Time    : 2019-06-14 19:45
// @Author  : victor！！
// @FileName: main.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	userPKG "shippy-micro/user-service/proto/user"
)

func main() {
	service := micro.NewService()
	service.Init()

	client := userPKG.NewUserServiceClient("go.micro.srv.user", service.Client())

	// 暂时将用户信息写死在代码中
	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	resp, err := client.Create(context.Background(), &userPKG.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("failed to create user error:%s\n", err)
	}
	if resp.Created {
		log.Println("create user success")
	}

	allResp, err := client.GetAll(context.Background(), &userPKG.Request{})
	if err != nil {
		log.Fatalf("failed to get user all error: %s\n", err)
	}
	for idx, user := range allResp.Users {
		log.Printf("user_%d: %s\n", idx, user)
	}

	authResp, err := client.Auth(context.Background(), &userPKG.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("failed to auth user error:%s\n", err)
	}
	log.Printf("token: %s\n", authResp.Token)
}
