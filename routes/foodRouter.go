package routes

import (
	"github.com/gin-gonic/gin"
	controller "restaurant-management/controllers"

)

func FoodRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/tables",controller.GetFoods())
	incomingRoutes.GET("/tables:table_id",controller.GetFood())
	incomingRoutes.POST("/tables",controller.CreateFood())
	incomingRoutes.PATCH("/tables/:table_id",controller.UpdateFood())
}