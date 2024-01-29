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

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func CreateOrder() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		bindJSONError := ginContext.BindJSON(&order)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusBadRequest,
				"CreateOrder():  Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(order)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateOrder(): Food struct validation failed",
				validationError.Error())
			return
		}

		var table models.Table
		findTableError := tableCollection.FindOne(context, bson.M{"table_id": order.Table_ID}).Decode(&table)
		if findTableError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateOrder(): Table fetched failed",
				findTableError.Error())
			return
		}

		createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Created_At = createdAt
		order.Updated_At = updatedAt

		order.ID = primitive.NewObjectID()
		order.Order_ID = order.ID.Hex()

		result, insertionError := orderCollection.InsertOne(context, &order)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateOrder(): Food insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func GetOrders() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, findError := orderCollection.Find(context, bson.M{})
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetOrders(): Orders fetched failed",
				findError.Error())
			return
		}

		defer cursor.Close(context)

		var orders []models.Order
		// var orders []bson.M
		cursorError := cursor.All(context, &orders)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetOrders(): Cursor iteration failed",
				cursorError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &orders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderID := ginContext.Param("order_id")
		var order models.Order

		findError := orderCollection.FindOne(context, bson.M{"order_id": orderID}).Decode(&order)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetOrder(): Order fetched failed",
				findError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &order)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		bindJSONError := ginContext.BindJSON(&order)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusBadRequest,
				"UpdateOrder():  Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		orderID := ginContext.Param("order_id")
		var updatedObj primitive.D
		var table models.Table

		if order.Table_ID != nil {
			findError := tableCollection.FindOne(context, bson.M{"table_id": order.Table_ID}).Decode(&table)
			if findError != nil {
				helpers.SendCustomErrorHelper(ginContext,
					http.StatusInternalServerError,
					"UpdateOrder(): Table fetched failed",
					findError.Error())
				return
			}

			updatedObj = append(updatedObj, bson.E{Key: "table_id", Value: order.Table_ID})
		}

		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: updatedAt})

		upsert := true
		filter := bson.M{"order_id": orderID}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, updateError := orderCollection.UpdateOne(
			context,
			filter,
			bson.D{{Key: "$set", Value: updatedObj}},
			&opt)

		if updateError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"UpdateOrder(): Error while updating order",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}

func OrderItemOrderCreator(order models.Order) string {
	context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_ID = order.ID.Hex()

	orderCollection.InsertOne(context, order)
	defer cancel()

	return order.Order_ID
}
