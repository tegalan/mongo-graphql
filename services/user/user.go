package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User entity
type User struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name  string             `json:"name,omitempty"`
	Email string             `json:"email,omitempty"`
}
