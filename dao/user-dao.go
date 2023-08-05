package dao

import (
	"auth-service/structs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

type UserDAO struct {
	Collection *mongo.Collection
	Logger     *slog.Logger
}

func NewUserDAO(client *mongo.Client, logger *slog.Logger) (*UserDAO, error) {
	return &UserDAO{
		Collection: client.Database("EEVEE").Collection("Users"),
		Logger:     logger,
	}, nil
}

func (userDAO *UserDAO) Add(ctx context.Context, req *structs.User) bool {
	_, err := userDAO.Collection.InsertOne(ctx, req)
	if err != nil {
		userDAO.Logger.Error(fmt.Sprintf("Failed to add user. %s", err))

		return false
	}

	return true
}

func (userDAO *UserDAO) Get(ctx context.Context, nickname string) *structs.User {
	var user *structs.User

	err := userDAO.Collection.FindOne(ctx, bson.D{{"nickname", nickname}}).Decode(&user)
	if err != nil {
		userDAO.Logger.Error(fmt.Sprintf("Failed to get user %s. %s", nickname, err))

		return nil
	}

	return user
}

func (userDAO *UserDAO) Delete(ctx context.Context, nickname string) bool {
	one, err := userDAO.Collection.DeleteOne(ctx, bson.D{{"nickname", nickname}})
	if err != nil {
		userDAO.Logger.Error(fmt.Sprintf("Failed to delete user %s. %s", nickname, err))

		return false
	}

	return one.DeletedCount > 0
}
