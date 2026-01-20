package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

var client *nats.Conn

func Connect(url string) {
	var err error
	client, err = nats.Connect(url)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	log.Println("Connected to NATS at", url)
}

func CreateStream(name string, subject string) nats.JetStreamContext {
	if client == nil || client.IsClosed() {
		log.Fatal("NATS client is not connected")
	}

	js, err := client.JetStream()
	if err != nil {
		log.Fatal("Failed to create JetStream context:", err)
	}
	
	stream, err := js.AddStream(&nats.StreamConfig{
		Name:      name,
		Subjects:  []string{subject + ".>"},
		Storage:   nats.FileStorage,
		Retention: nats.LimitsPolicy,
	})
	if err != nil {
		log.Fatal("Failed to add stream:", err)
	}

	log.Printf("Created NATS JetStream: %s (stream info: %+v)\n", name, stream)
	return js
}

func Close() {
	if client != nil && !client.IsClosed() {
		client.Close()
		log.Println("NATS connection closed")
	}
}
