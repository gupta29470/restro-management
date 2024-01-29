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

type InvoiceViewFormat struct {
	Invoice_ID       string
	Payment_Method   string
	Order_ID         string
	Payment_Status   *string
	Payment_Due      interface{}
	Table_Number     interface{}
	Payment_Due_Date time.Time
	Order_Details    interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func CreateInvoice() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		bindJSONError := ginContext.BindJSON(&invoice)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"UpdateInvoice(): Error while decoding invoice from request",
				bindJSONError.Error())
			return
		}

		validationError := validate.Struct(invoice)
		if validationError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"CreateFood(): Food struct validation failed",
				validationError.Error())
			return
		}

		var order models.Order
		findOrderError := orderCollection.FindOne(context, bson.M{"order_id": invoice.Order_ID}).Decode(&order)
		if findOrderError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusNotFound,
				"CreateInvoice(): Order not found",
				findOrderError.Error())
			return
		}

		invoice.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_ID = invoice.ID.Hex()
		invoice.Payment_Due_Date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))

		result, insertionError := invoiceCollection.InsertOne(context, &invoice)
		if insertionError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"CreateInvoice(): Invoice insertion failed",
				insertionError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &result)
	}
}

func GetInvoices() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, findError := invoiceCollection.Find(context, bson.M{})
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetInvoices(): Invoices fetched failed",
				findError.Error())
			return
		}

		defer cursor.Close(context)

		var invoices []models.Invoice
		cursorError := cursor.All(context, &invoices)
		if cursorError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetInvoices(): Cursor iteration failed",
				cursorError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, &invoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		invoiceID := ginContext.Param("invoice_id")
		var invoice models.Invoice

		findError := invoiceCollection.FindOne(context, bson.M{"invoice_id": invoiceID}).Decode(&invoice)
		if findError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetInvoices(): Invoices fetched failed",
				findError.Error())
			return
		}

		var invoiceView InvoiceViewFormat
		allOrders, allOrdersError := ItemsByOrder(invoice.Order_ID, ginContext)
		if allOrdersError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetInvoices(): Items by order error",
				findError.Error())
			return
		}

		if len(allOrders) <= 0 {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusInternalServerError,
				"GetInvoices(): No Orders found",
				"")
			return
		}

		invoiceView.Invoice_ID = invoice.Invoice_ID
		invoiceView.Order_ID = invoice.Order_ID
		invoiceView.Payment_Due_Date = invoice.Payment_Due_Date

		invoiceView.Payment_Method = "null"
		if invoice.Payment_Method != nil {
			invoiceView.Payment_Method = *invoice.Payment_Method
		}

		invoiceView.Payment_Status = *&invoice.Payment_Status
		invoiceView.Payment_Due = allOrders[0]["payment_due"]
		invoiceView.Table_Number = allOrders[0]["table_number"]
		invoiceView.Order_Details = allOrders[0]["order_items"]

		ginContext.JSON(http.StatusOK, &invoiceView)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		invoiceID := ginContext.Param("invoice_id")

		bindJSONError := ginContext.BindJSON(&invoice)
		if bindJSONError != nil {
			helpers.SendCustomErrorHelper(ginContext,
				http.StatusBadRequest,
				"UpdateInvoice(): Error while decoding invoice from request",
				bindJSONError.Error())
			return
		}

		var updateObj primitive.D

		if invoice.Payment_Method != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_method", Value: invoice.Payment_Method})
		}

		if invoice.Payment_Status != nil {
			updateObj = append(updateObj, bson.E{Key: "payment_status", Value: invoice.Payment_Status})
		}

		updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

		filter := bson.M{"invoice_id": invoiceID}
		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, updateError := invoiceCollection.UpdateOne(
			context,
			filter,

			bson.D{{Key: "$set", Value: updateObj}},
			&opt)

		if updateError != nil {
			helpers.SendCustomErrorHelper(
				ginContext,
				http.StatusInternalServerError,
				"UpdateInvoice(): Error while updating invoice",
				updateError.Error())
			return
		}

		ginContext.JSON(http.StatusOK, result)
	}
}
