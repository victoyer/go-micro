// @Time    : 2019-06-14 14:01
// @Author  : victor！！
// @FileName: handler.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	vessPKG "shippy-micro/vessel-service/proto/vessel"
)

type handler struct {
	repo Repository
}

func (this *handler) FindAvailable(ctx context.Context, req *vessPKG.Specification, rsp *vessPKG.Response) error {
	vessel, err := this.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	rsp.Vessel = vessel
	return nil
}

func (this *handler) Create(ctx context.Context, req *vessPKG.Vessel, rsp *vessPKG.Response) error {
	if err := this.repo.Create(req); err != nil {
		return err
	}
	rsp.Created = true
	rsp.Vessel = req
	return nil
}
