// @Time    : 2019-06-14 11:20
// @Author  : victor！！
// @FileName: main.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	consPKG "shippy-micro/consignment-service/proto/consignment"
	userPKG "shippy-micro/user-service/proto/user"
	vessPKG "shippy-micro/vessel-service/proto/vessel"
)

const (
	DB_NAME        = "consignment"
	CON_COLLECTION = "consignment"
	DB_HOST        = "127.0.0.1:27017"
	topic          = "vessel.service"
)

// AuthWrapper 是一个高阶函数， 入参传入的是 "下一步" 函数 出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-client 上下文中取出来的，再调用 user-service 将其做验证
// 认证通过则 fn() 函数继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// consignment-service 独立测试时不进行认证
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, rsp)
		}
		//  需要 user-service 进行认证
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]

		// Auth here
		authClient := userPKG.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		authResp, err := authClient.ValidateToken(context.Background(), &userPKG.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		log.Printf("Auth Resp: %s\n", authResp)
		return fn(ctx, req, rsp)
	}
}

func main() {
	service := micro.NewService(micro.Name("go.micro.srv.consignment"), micro.WrapHandler(AuthWrapper))
	service.Init()
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DB_HOST
	}
	sess, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("failed to create session error: %s\n", err)
	}
	vesselCli := vessPKG.NewVesselServiceClient("go.micro.srv.vessel", service.Client())
	// 创建nats服务发现实例
	//publisher := micro.NewPublisher(topic, service.Client())

	defer sess.Close()
	consPKG.RegisterConsignmentServiceHandler(service.Server(), &handler{Repository{sess}, vesselCli})

	if err := service.Run(); err != nil {
		log.Fatalf("failed to service run error: %s\n", err)
	}

}
