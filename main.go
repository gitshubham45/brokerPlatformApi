package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("stock broker platform")

	r := gin.Default()

	r.GET("health" , func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"health" : "ok"})
	})

	port := "8080"
	

	r.Run(":" + port)
}
