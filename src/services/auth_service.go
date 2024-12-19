package services

import (
	"context"
	"errors"
	"time"

	"register/src/config"
	"register/src/constants"
	"register/src/models"
	"register/src/models/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func Register(req schemas.RegisterSchema) error {
	client = config.MongoClient
	collection := client.Database(constants.DBNAME).Collection(constants.USER)

	filter := bson.M{constants.Email: req.Email}
	var existingUser models.User
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {
		return errors.New("user already registered")
	} else if err != mongo.ErrNoDocuments {
		return errors.New("database error")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	newUser := models.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}
