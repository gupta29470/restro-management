package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/database"
	"github.com/gupta29470/restro-mngmt/helpers"
	"github.com/gupta29470/restro-mngmt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func HashPassword(password string) string {
	encryptedPassword, encryptError := bcrypt.GenerateFromPassword([]byte(password), 14)
	if encryptError != nil {
		return ""
	}

	return string(encryptedPassword)
}

func VerifyPassword(password string, enteredPassword string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(password), []byte(enteredPassword))
	return compareError == nil
}

func Signup() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		bindJSONError := ginContext.BindJSON(&user)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"Signup(): Error while decoding food from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(user)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"Signup(): Food struct validation failed",
				validationError.Error())
			return
		}

		countEmail, countEmailError := userCollection.CountDocuments(context, bson.M{"email": user.Email})
		if countEmailError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"Signup(): Email checking error",
				countEmailError.Error())
			return
		}

		if countEmail > 0 {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"Signup(): User already signed up",
				"Email exists")
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		token, refreshToken := helpers.GenerateAllTokens(*user.Email, *user.First_Name, *user.Last_Name, *user.Phone_Number, user.User_ID)
		user.Token = &token
		user.RefreshToken = &refreshToken

		result, insertionError := userCollection.InsertOne(context, &user)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"Signup(): User insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func Login() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		bindJSONError := ginContext.BindJSON(&user)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"Login(): Error while decoding food from request",
				bindJSONError.Error())
			return
		}

		findError := userCollection.FindOne(context, bson.M{"email": user.Email}).Decode(&foundUser)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"User not found",
				findError.Error())
			return
		}

		isPasswordValid := VerifyPassword(*foundUser.Password, *user.Password)
		if !isPasswordValid {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"Password is wrong",
				"Password is wrong")
			return
		}

		token, refreshToken := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.Phone_Number, foundUser.User_ID)
		foundUser.Token = &token
		foundUser.RefreshToken = &refreshToken

		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_ID, ginContext)

		ginContext.JSON(http.StatusOK, &foundUser)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordsPerPageStr := ginContext.Query("records_per_page")
		pageNumberStr := ginContext.Query("page_number")

		recordsPerPage, recordsPerPageError := strconv.Atoi(recordsPerPageStr)
		pageNumber, pageNumberError := strconv.Atoi(pageNumberStr)

		if recordsPerPageError != nil || recordsPerPage < 1 {
			recordsPerPage = 1
		}

		if pageNumberError != nil || pageNumber < 1 {
			pageNumber = 1
		}

		startIndex := (pageNumber - 1) * recordsPerPage
		findOptions := options.Find()
		findOptions.SetSkip(int64(startIndex))
		findOptions.SetLimit(int64(recordsPerPage))

		cursor, findError := userCollection.Find(context, bson.M{}, findOptions)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetUsers(): Users fetched failed",
				findError.Error())
			return
		}

		defer cursor.Close(context)

		var users []models.User
		cursorError := cursor.All(context, &users)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetUsers(): Cursor iteration failed",
				cursorError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &users)

	}
}

func GetUser() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userID := ginContext.Param("user_id")
		var user models.User

		findError := userCollection.FindOne(context, bson.M{"user_id": userID}).Decode(&user)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"User not found",
				findError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &user)
	}
}
