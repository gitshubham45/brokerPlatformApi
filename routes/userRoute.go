package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/controllers"
	"github.com/gitshubham45/brokerPlatformApi/middlewares"
)

func UserRoute(api *gin.RouterGroup) {
	api.POST("/user/login", controllers.UserLogin)
	api.POST("/user/signup", controllers.UserSignup)
	api.GET("/user/access-token", middlewares.Authenticate, controllers.GetAccessToken)
}
