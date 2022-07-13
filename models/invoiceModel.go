package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	invoice_id       string             `json:"invoice_id"`
	order_id         string             `json:"order_id"`
	payment_method   *string            `json:"payment_method"validate:"eq=CARD|eq=CASH|eq="`
	payment_status   *string            `json:"payment_status"validate:"required,eq=PENDING|eq=PAID"`
	payment_due_date time.Time          `json:"payment_due_date"`
	created_at       time.Time          `json:"created_at"`
	updated_at       time.Time          `json:"updated_at"`
}
