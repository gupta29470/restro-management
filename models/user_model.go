package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_Name   *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name    *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email        *string            `json:"email" validate:"required"`
	Password     *string            `json:"password" validate:"required"`
	Phone_Number *string            `jsson:"phone_number" validate:"required"`
	Avatar       *string            `json:"avatar"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	Created_At   time.Time          `json:"created_at"`
	Updated_At   time.Time          `json:"updated_at"`
	User_ID      string             `json:"user_id"`
}
