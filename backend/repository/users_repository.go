package repository

import (
	"context"
	"errors"

	"github.com/chera-mihiretu/IKnow/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserById(ctx context.Context, userID string) (models.UserView, error)
	GetUserByEmail(ctx context.Context, email string) (models.UserView, error)
	GetListOfUsers(ctx context.Context, ids []primitive.ObjectID) ([]models.UserView, error)
}

type userRepository struct {
	users *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *userRepository {

	return &userRepository{
		users: db.Collection("users"),
	}
}

func (c *userRepository) GetUserById(ctx context.Context, userID string) (models.UserView, error) {
	var user models.UserView
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.UserView{}, errors.New("invalid user ID format")
	}
	err = c.users.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.UserView{}, err
	}
	return user, nil
}

func (c *userRepository) GetUserByEmail(ctx context.Context, email string) (models.UserView, error) {
	var user models.UserView
	err := c.users.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.UserView{}, err
	}
	return user, nil
}

func (c *userRepository) GetListOfUsers(ctx context.Context, ids []primitive.ObjectID) ([]models.UserView, error) {
	var users []models.UserView
	cursor, err := c.users.Find(ctx, bson.M{
		"_id": bson.M{"$in": ids},
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.UserView
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
