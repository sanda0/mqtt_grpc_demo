package grpc_services

import (
	"context"
	"time"

	"github.com/sanda0/mqtt_grpc_demo/internal/db"
	"github.com/sanda0/mqtt_grpc_demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SensorService struct {
	proto.UnimplementedSensorServiceServer
}

// GetSensorData implements SensorServiceInterface.
func (s *SensorService) GetSensor(ctx context.Context, in *proto.SensorRequest) (*proto.SensorResponse, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()
	data, exsits := db.SensorData[in.SensorId]
	if !exsits {
		return nil, status.Errorf(codes.NotFound, "Sensor not found")
	}

	return data, nil
}

// StreamSensorData implements SensorServiceInterface.
func (s *SensorService) StreamSensor(in *proto.SensorRequest, stream grpc.ServerStreamingServer[proto.SensorResponse]) error {
	for {
		db.Mu.Lock()
		data, exsits := db.SensorData[in.SensorId]
		db.Mu.Unlock()

		if !exsits {
			return status.Errorf(codes.NotFound, "Sensor not found")
		}

		if err := stream.Send(data); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
