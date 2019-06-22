build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-client .
run:
	docker run -e MICRO_REGISTRY=etcd  user-client
