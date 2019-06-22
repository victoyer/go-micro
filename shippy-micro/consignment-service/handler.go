// @Time    : 2019-06-14 11:24
// @Author  : victor！！
// @FileName: handler.go
// @Software: IntelliJ IDEA

package main

import (
	"context"
	consPKG "shippy-micro/consignment-service/proto/consignment"
	vessPKG "shippy-micro/vessel-service/proto/vessel"
)

type handler struct {
	repo Repository
	vesselClient vessPKG.VesselServiceClient
	//publisher micro.Publisher
}

func (this *handler) CreateConsignment(ctx context.Context, req *consPKG.Consignment, rsp *consPKG.Response) error {
	vReq := &vessPKG.Specification{
		Capacity:  int32(len(req.Container)),
		MaxWeight: req.Weight,
	}
	//if err := this.publisher.Publish(ctx, vReq); err != nil {
	//	return err
	//}
	vResp, err := this.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}
	req.VesselId = vResp.Vessel.Id
	consignment, err := this.repo.Create(req)
	if err != nil {
		return err
	}
	rsp.Consignment = consignment
	rsp.Created = true
	return nil
}

func (this *handler) GetConsignments(ctx context.Context, req *consPKG.Request, rsp *consPKG.Response) error {
	consignments, err := this.repo.GetAll()
	if err != nil {
		return err
	}
	rsp.Consignments = consignments
	return nil
}
