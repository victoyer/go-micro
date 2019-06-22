build:
	protoc -I. --go_out=plugins=micro:. ./proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .
run:
	docker run --net="host" -p 50053:50051 -e MICRO_REGISTRY=etcd -e DB_NAME="postgres" -e DB_HOST="127.0.0.1" -e DB_PORT="5432" -e DB_USER="postgres" -e DB_PASSWORD="postgres" -e MICRO_BROKER=nats -e MICRO_BROKER_ADDRESS=0.0.0.0:4222 user-service
