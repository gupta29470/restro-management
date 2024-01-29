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

type OrderItemPack struct {
	Table_ID    *string
	Order_Items []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func CreateOrderItems() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var orderItemPack OrderItemPack
		var order models.Order

		bindJSONError := ginContext.BindJSON(&orderItemPack)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateOrderItems(): Error while decoding order item from request",
				bindJSONError.Error())
			return
		}

		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		orderItemsToBeInserted := []interface{}{}
		order.Table_ID = orderItemPack.Table_ID
		order.Order_ID = OrderItemOrderCreator(order)

		for _, orderItem := range orderItemPack.Order_Items {
			orderItem.Order_ID = order.Order_ID

			validationError := validate.Struct(orderItem)
			if validationError != nil {
				helpers.SendCustomErrorHelper(ginContext,
					http.StatusBadRequest,
					"CreateOrderItems(): Food struct validation failed",
					validationError.Error())
				return
			}

			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			orderItem.Order_Item_ID = orderItem.ID.Hex()
			var num = toFixed(*orderItem.Unit_Price, 2)
			orderItem.Unit_Price = &num
			orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)
		}

		result, insertionError := orderItemCollection.InsertMany(context, orderItemsToBeInserted)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateOrderItems(): Food insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func GetOrderItems() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, findError := orderItemCollection.Find(context, bson.M{})
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetOrderItems(): Order Items fetched failed",
				findError.Error())
			return
		}

		defer cursor.Close(context)

		var orderItems []models.OrderItem
		cursorError := cursor.All(context, &orderItems)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetOrderItems(): Cursor iteration failed",
				cursorError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &orderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderItemID := ginContext.Param("order_item_id")
		var orderItem models.OrderItem

		findError := orderItemCollection.FindOne(context, bson.M{"order_item_id": orderItemID}).Decode(&orderItem)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"Order Item not found",
				findError.Error())
		}

		ginContext.JSON(http.StatusOK, &orderItem)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		orderItemID := ginContext.Param("order_id")
		allOrders, findError := ItemsByOrder(orderItemID, ginContext)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"Get Order Items By Order not found",
				findError.Error())
		}

		ginContext.JSON(http.StatusOK, &allOrders)
	}
}

func ItemsByOrder(orderItemID string, ginContext *gin.Context) (allOrders []primitive.M, findError error) {
	context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "order_id", Value: orderItemID}}}}

	lookupStageFood := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "food"},
			{Key: "localField", Value: "food_id"},
			{Key: "foreignField", Value: "food_id"},
			{Key: "as", Value: "food"}}}}
	unwindStageFood := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$food"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

	lookupStageOrder := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "order"},
			{Key: "localField", Value: "order_id"},
			{Key: "foreignField", Value: "order_id"},
			{Key: "as", Value: "order"}}}}
	unwindStageOrder := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$order"},
			{Key: "preserveNullAndEmptyArrays", Value: true}}}}

	lookupStageTable := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "table"},
			{Key: "localField", Value: "order.table_id"},
			{Key: "foreignField", Value: "table_id"},
			{Key: "as", Value: "table"}}}}
	unwindStageTable := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$table"},
			{Key: "preserveNullAndEmptyArrays", Value: true}}}}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "amount", Value: "$food.price"},
			{Key: "total_price_per_item", Value: bson.D{{Key: "$multiply", Value: bson.A{"$food.price", "$quantity"}}}},
			{Key: "total_count", Value: 1},
			{Key: "food_name", Value: "$food.name"},
			{Key: "food_image", Value: "$food.food_image"},
			{Key: "table_number", Value: "$table.table_number"},
			{Key: "table_id", Value: "$table.table_id"},
			{Key: "order_id", Value: "$order.order_id"},
			{Key: "price", Value: "$food.price"},
			{Key: "quantity", Value: 1}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "order_id", Value: "$order_id"},
				{Key: "table_id", Value: "$table_id"},
				{Key: "table_number", Value: "$table_number"}}},
			{Key: "payment_due", Value: bson.D{{Key: "$sum", Value: "$total_price_per_item"}}},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "order_items", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}

	projectStage2 := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "payment_due", Value: 1},
			{Key: "total_count", Value: 1},
			{Key: "table_number", Value: "$_id.table_number"},
			{Key: "order_items", Value: 1}}}}

	result, aggregateError := orderItemCollection.Aggregate(context, mongo.Pipeline{
		matchStage,
		lookupStageFood,
		unwindStageFood,
		lookupStageOrder,
		unwindStageOrder,
		lookupStageTable,
		unwindStageTable,
		projectStage,
		groupStage,
		projectStage2})

	if aggregateError != nil {
		helpers.SendCustomErrorHelper(
			ginContext,
			http.StatusBadRequest,
			"UpdateOrderItem():  Aggregation Error",
			aggregateError.Error())
		return
	}

	iterationError := result.All(context, &allOrders)
	if iterationError != nil {
		helpers.SendCustomErrorHelper(
			ginContext,
			http.StatusBadRequest,
			"UpdateOrderItem():  Iteration Error",
			iterationError.Error())
		return
	}

	return allOrders, iterationError

}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderItemID := ginContext.Param("order_item_id")

		var orderItem models.OrderItem
		bindJSONError := ginContext.BindJSON(&orderItem)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusBadRequest,
				"UpdateOrderItem():  Error while decoding menu from request",
				bindJSONError.Error())
			return
		}

		var updatedObj primitive.D

		if orderItem.Food_ID != nil {
			updatedObj = append(updatedObj, bson.E{Key: "food_id", Value: orderItem.Food_ID})
		}

		if orderItem.Quantity != nil {
			updatedObj = append(updatedObj, bson.E{Key: "quantity", Value: orderItem.Quantity})
		}

		if orderItem.Unit_Price != nil {
			updatedObj = append(updatedObj, bson.E{Key: "unit_price", Value: orderItem.Unit_Price})
		}

		orderItem.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updatedObj = append(updatedObj, bson.E{Key: "updated_at", Value: orderItem.Updated_At})

		upsert := true
		filter := bson.M{"order_item_id": orderItemID}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, updateError := orderItemCollection.UpdateOne(
			context,
			filter,
			bson.D{{Key: "$set", Value: updatedObj}},
			&opt)

		if updateError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"UpdateOrderItem(): Error while updating order item",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}
