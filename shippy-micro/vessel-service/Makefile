build:
	protoc -I. --go_out=plugins=micro:. ./proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .
run:
	docker run -p 50052:50051 -e MICRO_REGISTRY=mdns -e MICRO_SERVER_ADDRESS=:50052 vessel-service
