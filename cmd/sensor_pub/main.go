package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	broker := "tcp://localhost:1883"
	clientID := "sensor-publisher"

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connection Error: %v", token.Error())
	}

	log.Println("MQTT Publisher Connected!")

	sensors := []string{"sensors/room1", "sensors/room2", "sensors/room3"}

	for {
		for _, topic := range sensors {
			temp := 20.0 + rand.Float64()*10.0
			payload := fmt.Sprintf("%.2f", temp)

			token := client.Publish(topic, 0, false, payload)
			token.Wait()

			log.Printf("Published: %s -> %s°C", topic, payload)
		}
		time.Sleep(2 * time.Second)
	}
}
