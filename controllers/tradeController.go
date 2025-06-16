package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHoldings(c *gin.Context) {
	data , _ := os.ReadFile("mockData/holding.json")
	var holdings map[string]interface{}
	_ = json.Unmarshal(data,&holdings)

	c.JSON(http.StatusOK , gin.H{"holdings" : holdings })
}

func GetOrderBook(c *gin.Context) {
	data , _ := os.ReadFile("mockData/order.json")
	var orders map[string]interface{}
	_ = json.Unmarshal(data,&orders)
	
	c.JSON(http.StatusOK , gin.H{"orders" : orders})
}

func GetPositions(c *gin.Context) {
	data, _ := os.ReadFile("mockData/position.json")
	var positions map[string]interface{}
	_ = json.Unmarshal(data, &positions)

	c.JSON(http.StatusOK, positions)
}
