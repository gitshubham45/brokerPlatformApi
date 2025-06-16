package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHoldings(c *gin.Context) {
	data, err := os.ReadFile("mockData/holding.json")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to read mock file",
			"details": err.Error(),
		})
		return
	}

	var holdings map[string]interface{}
	if err := json.Unmarshal(data, &holdings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to parse JSON",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, holdings)
}

func GetOrderBook(c *gin.Context) {
	data, _ := os.ReadFile("mockData/order.json")
	var orders map[string]interface{}
	_ = json.Unmarshal(data, &orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetPositions(c *gin.Context) {
	data, _ := os.ReadFile("mockData/position.json")
	var positions map[string]interface{}
	_ = json.Unmarshal(data, &positions)

	c.JSON(http.StatusOK, positions)
}
