package events

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Version   string          `json:"version"`
	Source    string          `json:"source"`
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

type EventInterface interface {
	Subject() string
	Version() string
	Source() string
	Payload() any
}

type UserCreated struct {
	Firstname string `json:"firstname"`
}

func (u UserCreated) Subject() string { return "users.user.created" }
func (u UserCreated) Version() string { return "v1.0" }
func (u UserCreated) Source() string  { return "user-service" }
func (u UserCreated) Payload() any    { return u }
