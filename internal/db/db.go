package db

import (
	"sync"

	"github.com/sanda0/mqtt_grpc_demo/proto"
)

var (
	SensorData = make(map[string]*proto.SensorResponse)
	Mu         sync.Mutex
)
