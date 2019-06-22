// @Time    : 2019-06-14 11:32
// @Author  : victor！！
// @FileName: main.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"io/ioutil"
	"log"
	consPKG "shippy-micro/consignment-service/proto/consignment"
)

func parseFile(fileName string) (*consPKG.Consignment, error) {
	if fileName == "" {
		return nil, errors.New("parse file but file name is empty")
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *consPKG.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, err
	}
	return consignment, nil
}

func main() {
	service := micro.NewService()
	service.Init()

	client := consPKG.NewConsignmentServiceClient("go.micro.srv.consignment", service.Client())

	consignment, err := parseFile("consignment.json")
	if err != nil {
		log.Fatalf("failed to parse file error: %s\n", err)
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiOWI3Yzk2NjAtZGVjOC00YzUyLWFhYmItYTY5ZWYyMmQxOWZkIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoiZXdhbi52YWxlbnRpbmU4OUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCRJMExMTExweHlEU05jZTV6OWJvdUZPU3JSaDZxYWd2SXpBMTZCbTBqeXdNdnZhWFV2SFpDbSJ9LCJleHAiOjE1NjA4Mjg4MzMsImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.fmXtDczrIfPYWcxm_Jzlu01lR7UWRiNobEzKY50ao44"
	// 创建带有用户 token 的 context
	// consignment-service 服务端将从中取出 token，解密取出用户身份
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("failed to create consignment error: %s\n", err)
	}
	if resp.Created {
		log.Println("create consignment success")
	}

	allResp, err := client.GetConsignments(tokenContext, &consPKG.Request{})
	if err != nil {
		log.Fatalf("failed to get consignment all error:%s\n", err)
	}
	for _, val := range allResp.Consignments {
		log.Println(val)
	}
}
