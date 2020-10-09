package post

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service interface
type Service interface {
	Create(*Post) error
	All() ([]*Post, error)
	FindByUser(string) ([]*Post, error)
}

// MongoService struct
type MongoService struct {
	db *mongo.Database
}

// NewMongoService instance
func NewMongoService(db *mongo.Database) Service {
	return &MongoService{db}
}

// FindByUser func
func (s *MongoService) FindByUser(uid string) ([]*Post, error) {
	id, _ := primitive.ObjectIDFromHex(uid)
	res, err := s.db.Collection("posts").Find(context.TODO(), bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}
	var posts []*Post
	for res.Next(context.TODO()) {
		var p Post
		res.Decode(&p)

		posts = append(posts, &p)
	}

	return posts, nil

}

// All post
func (s *MongoService) All() ([]*Post, error) {
	res, err := s.db.Collection("posts").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var posts []*Post
	for res.Next(context.TODO()) {
		var p Post
		res.Decode(&p)

		posts = append(posts, &p)
	}

	return posts, nil
}

// Create Post
func (s *MongoService) Create(p *Post) error {
	p.ID = primitive.NewObjectID()
	_, err := s.db.Collection("posts").InsertOne(context.TODO(), p)
	return err
}
