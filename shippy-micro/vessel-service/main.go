// @Time    : 2019-06-14 13:52
// @Author  : victor！！
// @FileName: main.go
// @Software: IntelliJ IDEA

package main

import (
	"github.com/micro/go-micro"
	"log"
	"os"
	vessPKG "shippy-micro/vessel-service/proto/vessel"
)

const (
	DB_NAME        = "vessel"
	CON_COLLECTION = "vessel"
	DB_HOST        = "127.0.0.1:27017"
)

func CreateDummyData(repo *Repository) {
	defer repo.Close()
	vessels := []*vessPKG.Vessel{
		{Id: "vessel_001", Name: "victor", MaxWeight: 2000000, Capacity: 50000},
	}
	for _, val := range vessels {
		if err := repo.Create(val); err != nil {
			log.Fatalf("failed to create vessel error: %s\n", err)
		}
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DB_HOST
	}

	sess, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("failed craete mongo session error: %s\n", err)
	}
	defer sess.Close()

	CreateDummyData(&Repository{sess: sess.Copy()})

	service := micro.NewService(micro.Name("go.micro.srv.vessel"))
	service.Init()

	vessPKG.RegisterVesselServiceHandler(service.Server(), &handler{Repository{sess: sess}})

	if err := service.Run(); err != nil {
		log.Fatalf("failed to service run error: %s\n", err)
	}
}
