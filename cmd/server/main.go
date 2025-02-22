package main

import (
	"log"
	"net"

	"github.com/sanda0/mqtt_grpc_demo/internal/grpc_services"
	mqttservices "github.com/sanda0/mqtt_grpc_demo/internal/mqtt_services"
	"github.com/sanda0/mqtt_grpc_demo/proto"
	"google.golang.org/grpc"
)

func main() {

	mqttClient := mqttservices.ConnectMQTT("tcp://localhost:1883")
	defer mqttClient.Disconnect(250)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSensorServiceServer(grpcServer, &grpc_services.SensorService{})
	log.Println("Starting server on port :9090")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
