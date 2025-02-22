package mqttservices

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sanda0/mqtt_grpc_demo/internal/db"
	"github.com/sanda0/mqtt_grpc_demo/proto"
)

func ConnectMQTT(broker string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connection Error: %v", token.Error())
	}

	topic := "sensors/+"
	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Subscription Error: %v", token.Error())
	}

	log.Println("Connected to MQTT and subscribed to sensors")
	return client
}

func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	sensorID := msg.Topic()
	temperature := string(msg.Payload())

	log.Printf("Received MQTT data: %s -> %s°C", sensorID, temperature)

	db.SensorData[sensorID] = &proto.SensorResponse{
		SensorId: sensorID,
		Temper:   parseFloat(temperature),
		Time:     time.Now().Format(time.RFC3339),
	}
}

func parseFloat(value string) float64 {
	var temp float64
	fmt.Sscanf(value, "%f", &temp)
	return temp
}
