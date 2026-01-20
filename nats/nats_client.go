package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

var client *nats.Conn

func Connect() {
	natsClient, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer natsClient.Close()
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
