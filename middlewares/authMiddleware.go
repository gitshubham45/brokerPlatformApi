package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/helpers"
	"github.com/golang-jwt/jwt"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("Authorization")
	if clientToken == "" {
		fmt.Println("No Authorization header provided")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
		c.Abort()
		return
	}

	if !strings.HasPrefix(clientToken, "Bearer ") {
		fmt.Println("Invalid Authorization header format")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		c.Abort()
		return
	}

	accessToken := strings.TrimPrefix(clientToken, "Bearer ")

	token, err := helpers.ValidateTokens(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	fmt.Printf("token : %s", token)

	claims, ok := token.Claims.(jwt.MapClaims)

	fmt.Printf("claims : %s", claims)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	expiry, ok := claims["exp"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid expiration time in token"})
		c.Abort()
		return
	}

	if int64(expiry) < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token expired"})
		c.Abort()
		return
	}

	c.Set("email", claims["email"])

	c.Next()
}
