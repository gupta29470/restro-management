package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *int               `json:"quantity" validate:"required"`
	Unit_Price    *float64           `json:"unit_price" validate:"required"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	Food_ID       *string            `json:"food_id" validate:"required"`
	Order_Item_ID string             `json:"order_item_id"`
	Order_ID      string             `json:"order_id" validate:"required"`
}
