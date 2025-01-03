package repository

import (
	"context"
	"errors"

	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthMongo struct {
	db *MongoDB
}

func NewAuthMongo(db *MongoDB) *AuthMongo {
	return &AuthMongo{
		db: db,
	}
}

func (r *AuthMongo) CreateUser(user models.UserModel) (string, error) {
	userCollection := r.db.database.Collection("users")
	var existingUser models.UserModel
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return "", errors.New("user with this email already exists")
	}
	if err != mongo.ErrNoDocuments {
		return "", err
	}

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (r *AuthMongo) IsUserValid(email, password string) (string, error) {
	userCollection := r.db.database.Collection("users")
	var existingUser models.UserModel

	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&existingUser)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password_hash), []byte(password)); err != nil {
		return "", err
	}

	return existingUser.ID.Hex(), nil
}
