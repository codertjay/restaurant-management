package controller

import (
	"context"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food
		err :=foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil{
			c.Json(http.StatusInternalServerError,gin.M{"error":"Error occured fetching the food item"})
		}
		c.Json(http.StatusOk,food)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {

}

func toFixed(num float64, precision int) float64 {

}
