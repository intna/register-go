package services

import (
	"context"
	"time"

	"register/src/config"
	"register/src/constants"
	"register/src/models"
	"register/src/models/schemas"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func Register(c *gin.Context, req schemas.RegisterSchema) {

	client = config.MongoClient
	collection := client.Database(constants.DBNAME).Collection(constants.USER)

	filter := bson.M{constants.Email: req.Email}
	var existingUser models.User
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {
		c.JSON(constants.INTERNAL_SERVER_ERROR, gin.H{"error": "User already registered"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(constants.INTERNAL_SERVER_ERROR, gin.H{"error": "Database error"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(constants.INTERNAL_SERVER_ERROR, gin.H{"error": "Failed to hash password"})
		return
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
		c.JSON(constants.INTERNAL_SERVER_ERROR, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(constants.OK, gin.H{
		"message":       "Registration successful",
		constants.Email: req.Email,
	})
}
