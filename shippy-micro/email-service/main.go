// @Time    : 2019-06-15 11:09
// @Author  : victor！！
// @FileName: main.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/nats"
	"log"
	userPKG "shippy-micro/user-service/proto/user"
)

const (
	topic = "user.created"
)

type Subscriber struct{}

func (this *Subscriber) Process(ctx context.Context, user *userPKG.User) error {
	log.Println("[Sending email to]:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(micro.Name("go.micro.srv.email"))
	srv.Init()

	pubSub := srv.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connection error: %s\n", err)
	}

	if err := micro.RegisterSubscriber(topic, srv.Server(), Subscriber{}); err != nil {
		log.Fatalf("failed to register subScribe error: %s\n", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("srv run error: %s\n", err)
	}
}
