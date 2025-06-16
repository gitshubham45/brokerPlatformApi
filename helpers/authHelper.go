package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gitshubham45/brokerPlatformApi/db"
	"github.com/gitshubham45/brokerPlatformApi/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

var userCollection *mongo.Collection = db.OpenCollection(db.Client , "user")

type SignedDetails struct {
	Email string
	ID    string
	jwt.StandardClaims
}

func GenerateTokens(user models.User) (string, string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id" : user.ID,
		"exp":   time.Now().Local().Add(10 * time.Minute).Unix(),
	})


	// claims := SignedDetails{
	// 	Email: user.Email,
	// 	ID:    user.ID,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Local().Add((time.Minute * time.Duration(10))).Unix(),
	// 	},
	// }

	// refreshClaims := SignedDetails{
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Local().Add((time.Hour * time.Duration(24))).Unix(),
	// 	},
	// }

	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp" : time.Now().Local().Add(24 * time.Hour).Unix(),
	})

	// token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	token, _ := claims.SignedString([]byte(SECRET_KEY))
	// refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := refreshClaims.SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	return token, refreshToken, nil
}

func UpdateTokens(signedAccessToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedAccessToken})
	updateObj = append(updateObj, bson.E{Key: "refreshToken", Value: signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updatedAt", Value: updatedAt})

	upsert := true
	filter := bson.M{"_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key : "$set" , Value : updateObj},
		},
		&opt,
	)

	if err != nil {
		fmt.Printf("Error updating tokens - %s", err)
	}
}

func ValidateTokens(accessToken string) (*jwt.Token , error ){
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
