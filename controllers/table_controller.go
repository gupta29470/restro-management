package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gupta29470/restro-mngmt/database"
	"github.com/gupta29470/restro-mngmt/helpers"
	"github.com/gupta29470/restro-mngmt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func CreateTable() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var table models.Table

		bindJSONError := ginContext.BindJSON(&table)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateTable(): Error while decoding food from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(table)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateTable(): Food struct validation failed",
				validationError.Error())
			return
		}

		table.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		table.ID = primitive.NewObjectID()
		table.Table_ID = table.ID.Hex()

		result, insertionError := tableCollection.InsertOne(ctx, table)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateTable(): Table insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}

func GetTables() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, findError := tableCollection.Find(ctx, bson.M{})
		if findError != nil {
			if findError != nil {
				helpers.SendCustomErrorHelper(ginContext,
					http.StatusInternalServerError,
					"GetTables(): Tables fetched failed",
					findError.Error())
				return
			}
		}
		var allTables []bson.M
		cursorError := cursor.All(ctx, &allTables)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetTables(): Cursor iteration failed",
				cursorError.Error())
			return
		}
		ginContext.JSON(http.StatusOK, allTables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		tableId := ginContext.Param("table_id")
		var table models.Table

		findError := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"Table not found",
				findError.Error())
		}

		ginContext.JSON(http.StatusOK, &table)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var table models.Table

		tableId := ginContext.Param("table_id")

		bindJSONError := ginContext.BindJSON(&table)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusBadRequest,
				"UpdateTable():  Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		var updateObj primitive.D

		if table.Number_Of_Guests != nil {
			updateObj = append(updateObj, bson.E{Key: "number_of_guests", Value: table.Number_Of_Guests})
		}

		if table.Table_Number != nil {
			updateObj = append(updateObj, bson.E{Key: "table_number", Value: table.Table_Number})
		}

		table.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		filter := bson.M{"table_id": tableId}

		result, updateError := tableCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{Key: "$set", Value: updateObj}},
			&opt,
		)

		if updateError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"UpdateTable(): Error while updating table",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}
