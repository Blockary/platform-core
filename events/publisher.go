package events

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func PublishEvent(js nats.JetStreamContext, e EventInterface) error {
	payload, err := json.Marshal(e.Payload())
	if err != nil {
		return err
	}

	event := Event{
		ID:        uuid.NewString(),
		Type:      e.Subject(),
		Version:   e.Version(),
		Source:    e.Source(),
		Timestamp: time.Now().UTC(),
		Data:      payload,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, err = js.Publish(e.Subject(), eventBytes)
	if err != nil {
		return err
	}

	log.Printf("Published event %s id=%s", e.Subject(), event.ID)
	return nil
}
