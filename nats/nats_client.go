package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

var client *nats.Conn

func Connect(url string) {
	var err error
	client, err = nats.Connect(url) // assign to package-level variable
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}
	log.Println("Connected to NATS at", url)
}

func CreateSteam(name string, subject string) nats.JetStreamContext {
	js, _ := client.JetStream()

	js.AddStream(&nats.StreamConfig{
		Name:      name,
		Subjects:  []string{subject + ".>"},
		Storage:   nats.FileStorage,
		Retention: nats.LimitsPolicy,
	})

	return js
}

func Close() {
	if client != nil && !client.IsClosed() {
		client.Close()
		log.Println("NATS connection closed")
	}
}
