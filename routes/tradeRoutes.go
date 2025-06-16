package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/controllers"
	"github.com/gitshubham45/brokerPlatformApi/middlewares"
)

func TradeRoutes(api *gin.RouterGroup) {
	api.Use(middlewares.Authenticate)
	{
		api.GET("/holdings", controllers.GetHoldings)
		api.GET("/orderbook", controllers.GetOrderBook)
		api.GET("/positions", controllers.GetPositions)
	}
}
