package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Departments defines the database model for the departments table
type Departments struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	SchoolID    primitive.ObjectID `bson:"school_id" json:"school_id"`
	Years       int                `bson:"years" json:"years"`
	CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
