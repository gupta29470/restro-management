package helpers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gupta29470/restro-mngmt/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

type SignedDetails struct {
	Email        string
	First_Name   string
	Last_Name    string
	Phone_Number string
	User_ID      string
	jwt.StandardClaims
}

const (
	secretKey = "HwfoReq9lKSvLm+ONuAQ21JFkMeFx3JxqngToE48D4s="
)

func GenerateAllTokens(email string, firstName string, lastName string, phoneNumber string, userID string) (string, string) {
	claims := &SignedDetails{
		Email:        email,
		First_Name:   firstName,
		Last_Name:    lastName,
		Phone_Number: phoneNumber,
		User_ID:      userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(124)).Unix(),
		},
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, tokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	refreshToken, refreshTokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if tokenError != nil || refreshTokenError != nil {
		log.Panic(tokenError, refreshTokenError)
	}

	return token, refreshToken
}

func UpdateAllTokens(token string, refreshToken string, userID string, ginContext *gin.Context) {
	context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: token})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: refreshToken})
	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

	upsert := true
	filter := bson.M{"user_id": userID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, updateError := userCollection.UpdateOne(context, filter,
		bson.D{{Key: "$set", Value: updateObj}},
		&opt)

	if updateError != nil {
		SendCustomErrorHelper(
			ginContext,
			http.StatusInternalServerError,
			"Update Toekn(): Error while updating tokens",
			updateError.Error())
		return
	}
}

func ValidateTokens(clientToken string, ginContext *gin.Context) (claims *SignedDetails) {
	token, tokenParseError := jwt.ParseWithClaims(
		clientToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if tokenParseError != nil {
		SendCustomErrorHelper(
			ginContext,
			http.StatusInternalServerError,
			"Validate Tokens(): Error while parsing tokens",
			tokenParseError.Error())
		return
	}

	claims, claimsError := token.Claims.(*SignedDetails)
	if !claimsError {
		SendCustomErrorHelper(
			ginContext,
			http.StatusInternalServerError,
			"Validate Tokens(): Claims Invalid",
			"")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		SendCustomErrorHelper(
			ginContext,
			http.StatusInternalServerError,
			"Validate Tokens(): Token Expired",
			"Token Expired")
		return
	}

	return claims

}
