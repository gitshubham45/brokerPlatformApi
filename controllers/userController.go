package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gitshubham45/brokerPlatformApi/db"
	"github.com/gitshubham45/brokerPlatformApi/helpers"
	"github.com/gitshubham45/brokerPlatformApi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")

func getPasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(hashedPassword)
}

func VerifyPassword(userPassword string, hashedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or passwor is incorrect")
		check = false
	}
	return check, msg
}

func UserSignup(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials - " + err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	//check if user is in already in db
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": req.Email})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "user is already there"})
		return
	}

	// if not create password hash
	passwordHash := getPasswordHash(req.Password)

	newUser := models.User{
		ID:           primitive.NewObjectID().Hex(),
		Email:        req.Email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// generate access token 
	// generate refresh token
	accessToken , refreshToken , _ := helpers.GenerateTokens(newUser)

	newUser.RefreshToken = refreshToken
	newUser.AccessToken = accessToken

	// save user 

	insertOneResult , err := userCollection.InsertOne(ctx , newUser) 


	if err != nil {
		msg := fmt.Sprintf("User not created : %s ", err)
		c.JSON(http.StatusInternalServerError , gin.H{"error" : msg})
	}

	c.JSON(http.StatusOK , gin.H{
		"message" : "User created successfully",
		"user" : map[string]interface{}{
			"_id" : newUser.ID,
			"email" : newUser.Email,
			"accessToken" : accessToken,
			"refreshToken" : refreshToken,
		},
		"insertOneResult" : insertOneResult,
	})
}



func UserLogin(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials - " + err.Error()})
		return
	}

	ctx , cancel := context.WithTimeout(context.Background() , 100 * time.Second)
	defer cancel()

	var foundUser models.User
	err := userCollection.FindOne(ctx , bson.M{"email" : req.Email}).Decode(&foundUser)
	if err != nil {
		fmt.Printf("error in finding user %s \n", err)
		c.JSON(http.StatusInternalServerError , gin.H{"error" : "email or password id incorrect"})
		return
	}

	isValidPassword , msg := VerifyPassword(req.Password , foundUser.PasswordHash)
	if !isValidPassword{
		fmt.Printf("Error validating password - %s \n" , err)
		c.JSON(http.StatusInternalServerError , gin.H{"error" : msg})
		return
	}

	// generate tokens
	accessToken , refreshToken , _ := helpers.GenerateTokens(foundUser)


	// update tokens
	helpers.UpdateTokens(accessToken , refreshToken , foundUser.ID)

	err = userCollection.FindOne(ctx , bson.M{"_id" : foundUser.ID}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : err})
		return
	}

	c.JSON(http.StatusOK , gin.H{
		"message" : "user login successful",
		"user" : foundUser,
	})
}
