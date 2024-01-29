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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func CreateMenu() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var menu models.Menu
		bindJSONError := ginContext.BindJSON(&menu)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateMenu(): Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(menu)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateMenu(): Menu struct validation failed",
				validationError.Error())
			return
		}

		menu.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_ID = menu.ID.Hex()

		result, insertError := menuCollection.InsertOne(ctx, &menu)
		if insertError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateMenu(): Menu insertion failed",
				insertError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func GetMenus() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		menus, findError := menuCollection.Find(context.TODO(), bson.M{})
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetMenus(): Error while fetching menu from DB",
				findError.Error())
			return
		}

		var allMenus []bson.M
		decodeError := menus.All(ctx, &allMenus)
		if decodeError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateFood(): Error while decoding menu",
				decodeError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		menuID := ginContext.Param("menu_id")
		var menu models.Menu

		decodeError := menuCollection.FindOne(ctx, bson.M{"menu_id": menuID}).Decode(&menu)
		if decodeError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"GetMenu(): Menu not found",
				decodeError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &menu)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		bindJSONError := ginContext.BindJSON(&menu)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"UpdateMenu(): Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		menuID := ginContext.Param("menu_id")
		filter := bson.M{"menu_id": menuID}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil && !inTimeSpan(menu.Start_Date, menu.End_Date) {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"UpdateMenu(): Start or End date is invalid. Kindly update the time",
				"")
			return
		}

		updateObj = append(updateObj, bson.E{Key: "start_date", Value: menu.Start_Date})
		updateObj = append(updateObj, bson.E{Key: "end_date", Value: menu.End_Date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{Key: "name", Value: menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{Key: "category", Value: menu.Category})
		}

		updatedMenuTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedMenuTime})

		upsert := true

		options := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, updateError := menuCollection.UpdateOne(context,
			filter,
			bson.D{{Key: "$set", Value: updateObj}},
			&options)

		if updateError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"UpdateMenu(): Update menu failed",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func inTimeSpan(start *time.Time, end *time.Time) bool {
	return start.After(time.Now()) && end.After(*start)
}
