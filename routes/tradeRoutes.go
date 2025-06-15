package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/middlewares"
)

func TradeRoutes(api *gin.RouterGroup) {
	api.Use(middlewares.Authenticate)
	{
		api.GET("/holdings", GetHoldings)
		api.GET("/orderbook", GetOrderBook)
		api.GET("/positions", GetPositions)
	}
}
