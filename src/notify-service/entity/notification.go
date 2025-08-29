package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationEntity struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	UserID    string                 `bson:"user_id" json:"user_id"`
	Title     string                 `bson:"title,omitempty" json:"title,omitempty"`
	Body      string                 `bson:"body" json:"body"`
	Channel   string                 `bson:"channel,omitempty" json:"channel,omitempty"`
	Payload   map[string]interface{} `bson:"payload,omitempty" json:"payload,omitempty"`
	IsRead    bool                   `bson:"is_read" json:"is_read"`
	CreatedAt time.Time              `bson:"created_at" json:"created_at"`
	ReadAt    *time.Time             `bson:"read_at,omitempty" json:"read_at,omitempty"`
	Meta      map[string]any         `bson:"meta,omitempty" json:"meta,omitempty"`
}
