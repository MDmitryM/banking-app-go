package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Email         string             `bson:"email"`
	Username      string             `bson:"username"`
	Password_hash string             `bson:"password_hash"`
	CreatedAt     time.Time          `bosn:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}
