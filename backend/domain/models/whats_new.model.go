package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// whats_new.model.go defines the database model for whats_new table
type WhatsNew struct {
	ID        primitive.ObjectID `bson:"id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
