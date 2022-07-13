package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"validate:"required"`
	Category   string             `json:"category"validate:"required"`
	Star_Date  *time.Time         `json:"star_date"`
	End_Date   *time.Time         `json:"end_date"`
	created_at time.Time          `json:"created_at"`
	updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"food_id"validate:"required"`
}
