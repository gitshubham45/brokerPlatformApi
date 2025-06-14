package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/controllers"
)

func UserRoute(api *gin.RouterGroup){
	api.POST("/user/login" , controllers.UserLogin)
	api.POST("/user/signup", controllers.UserSignup)
}