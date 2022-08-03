package controller

import (
	"context"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordperPage, err := strconv.Atoi("recordPerPage")
		if err != nil || recordperPage < 1 {
			recordperPage = 10
		}
		page,err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 10
		}
		startIndex := (page-1) * recordperPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{
			{"$match", bson.D{{}}},
		}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id",0},
			{"total_count",1},
			{"user_items",bson.D{{"$slice",[]interface{}{"$data",startIndex,recordperPage}}}}
			}}
		}
		result,err:= Aggregator(ctx,mongo.Pipeline{
			matchStage,
			projectStage,
		})
		defer cancel()
		if err!=nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error":"Error occured while listing user items"})
		}

		var allUsers []bson.M
		if err = result.All(ctx,&allUsers); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK,allUsers)

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("user_id")
		var user models.User
		 err :=userCollection.FindOne(ctx,bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError,
				gin.H{"error":"Error occured while listing user items"})
		}
		c.JSON(http.StatusOK,user)
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
