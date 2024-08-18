package models

import (
	"gorm.io/gorm"
)

type KafkaEvent struct {
	gorm.Model
	EventType string
	Payload   []byte
	Processed bool
}