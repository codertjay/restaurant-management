package main

import(
	"os"
	"github.com/gin-gonic/gin"
	"restaurant-management/database"
	"restaurant-management/routes"
	"restaurant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)
var foodCollection *mongo.Collection = database.OpenCollection(
	database.Client,"food")

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	
	router.Run(":"+port)


}