package models

import (
	"gorm.io/gorm"
)

type KafkaEvent struct {
    gorm.Model
    EventType string
    Payload   string `gorm:"type:text"`
    Processed bool
}