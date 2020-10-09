package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service struct
type Service interface {
	FindByID(string) (*User, error)
	All() ([]*User, error)
	Create(*User) error
}

// MongoService struct
type MongoService struct {
	db *mongo.Database
}

// NewUserService func
func NewUserService(db *mongo.Database) Service {
	return &MongoService{db}
}

// FindByID implementation
func (s *MongoService) FindByID(id string) (*User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("FindByID error parse id: %s ->%s\n", id, err.Error())
		return nil, err
	}

	var u User
	if err := s.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&u); err != nil {
		fmt.Printf("FindByID errr: %s %v \n", err.Error(), bson.M{"_id": objID})
		return nil, err
	}
	return &u, nil
}

// All user impl
func (s *MongoService) All() ([]*User, error) {
	res, err := s.db.Collection("users").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*User
	for res.Next(context.TODO()) {
		var user User
		res.Decode(&user)

		users = append(users, &user)
	}

	return users, nil
}

// Create implementation
func (s *MongoService) Create(u *User) error {
	u.ID = primitive.NewObjectID()
	_, err := s.db.Collection("users").InsertOne(context.TODO(), u)
	return err
}
