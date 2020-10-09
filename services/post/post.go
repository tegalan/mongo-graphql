package post

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post struct
type Post struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Body   string             `json:"body,omitempty"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id,omitempty"`
}
