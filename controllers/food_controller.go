package controllers

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gupta29470/restro-mngmt/database"
	"github.com/gupta29470/restro-mngmt/helpers"
	"github.com/gupta29470/restro-mngmt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(price float64, precision int) float64 {
	power := math.Pow(10, float64(precision))
	return float64(round(price*power)) / power
}

func CreateFood() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		var menu models.Menu

		bindJSONError := ginContext.BindJSON(&food)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateFood(): Error while decoding food from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(food)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateFood(): Food struct validation failed",
				validationError.Error())
			return
		}

		findMenuError := menuCollection.FindOne(context, bson.M{"menu_id": food.Menu_ID}).Decode(&menu)
		if findMenuError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"CreateFood(): Menu not found",
				findMenuError.Error())
			return
		}

		food.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_ID = food.ID.Hex()
		price := toFixed(*food.Price, 2)
		food.Price = &price

		result, insertionError := foodCollection.InsertOne(context, &food)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateFood(): Food insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func GetFoods() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordsPerPageStr := ginContext.Query("records_per_page")
		pageNumberStr := ginContext.Query("page_number")

		recordsPerPage, recordsPerPageError := strconv.Atoi(recordsPerPageStr)
		pageNumber, pageNumberError := strconv.Atoi(pageNumberStr)

		if recordsPerPageError != nil || recordsPerPage < 1 {
			recordsPerPage = 10
		}

		if pageNumberError != nil || pageNumber < 1 {
			pageNumber = 1
		}

		startIndex := (pageNumber - 1) * recordsPerPage
		findOptions := options.Find()
		findOptions.SetSkip(int64(startIndex))
		findOptions.SetLimit(int64(recordsPerPage))

		cursor, findError := foodCollection.Find(context, bson.M{}, findOptions)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetFoods(): Foods fetched failed",
				findError.Error())
			return
		}

		defer cursor.Close(context)

		var foods []models.Food
		// var foods []bson.M

		cursorError := cursor.All(context, &foods)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetFoods(): Cursor iteration failed",
				cursorError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &foods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := ginContext.Param("food_id")

		var food models.Food

		findError := foodCollection.FindOne(context, bson.M{"food_id": foodId}).Decode(&food)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"Food not found",
				findError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, food)
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		bindJSONError := ginContext.BindJSON(&food)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusBadRequest,
				"UpdateFood():  Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		foodID := ginContext.Param("food_id")
		var updateObj primitive.D

		if food.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: food.Name})
		}

		if food.Price != nil {
			updateObj = append(updateObj, bson.E{Key: "price", Value: food.Price})
		}

		if food.Food_Image != nil {
			updateObj = append(updateObj, bson.E{Key: "food_image", Value: food.Food_Image})
		}

		if food.Menu_ID != nil {
			var menu models.Menu

			decodeError := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_ID}).Decode(&menu)
			if decodeError != nil {
				helpers.SendCustomErrorHelper(
					ginContext,
					http.StatusNotFound,
					"UpdateFood(): Menu not found",
					decodeError.Error())
				return
			}

			updateObj = append(updateObj, bson.E{Key: "menu_id", Value: food.Menu_ID})
		}

		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		filter := bson.M{"food_id": foodID}

		result, updateError := foodCollection.UpdateOne(ctx,
			filter,
			bson.D{{Key: "$set", Value: updateObj}},
			&opt)

		if updateError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"UpdateFood(): Error while updating food",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}
