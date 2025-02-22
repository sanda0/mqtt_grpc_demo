protoc:
	protoc --go_out=. --go-grpc_out=. api/sensor.proto

runserver:
	go run cmd/server/main.go

runpublisher:
	go run cmd/sensor_pub/main.go