package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/db"
	"github.com/gitshubham45/brokerPlatformApi/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file %s \n" ,err)
	}
	fmt.Println("stock broker platform")

	db.Init()

	r := gin.Default()

	r.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	api := r.Group("/api")

	routes.UserRoute(api)
	routes.TradeRoutes(api)

	port := "8080"

	port = os.Getenv("PORT")

	r.Run(":" + port)
}
