package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(usersCollection *mongo.Collection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		type RegisterBody struct {
			Username string
			Password string
		}

		var registerRequest RegisterBody
		if err := ctx.BindJSON(&registerRequest); err != nil {
			panic(err)
		}

		// TODO: Only insert if username unique

		hash, hashErr := HashPassword(registerRequest.Password)
		if hashErr != nil {
			panic(hashErr)
		}
		registerRequest.Password = hash

		_, err := usersCollection.InsertOne(context.TODO(), registerRequest)
		if err == nil {
			ctx.JSON(200, gin.H{})
		}
		if err != nil {
			panic(err)
		}
	}
}

func Login(usersCollection *mongo.Collection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		type LoginBody struct {
			Username string
			Password string
		}

		var loginBody LoginBody
		if err := ctx.BindJSON(&loginBody); err != nil {
			panic(err)
		}

		// TODO: Check if no user found
		var user LoginBody
		if err := usersCollection.FindOne(context.TODO(), bson.M{"username": loginBody.Username}).Decode(&user); err != nil {
			panic(err)
		}

		hash := user.Password
		if passed := CheckPasswordHash(loginBody.Password, hash); passed {
			ctx.JSON(200, gin.H{"message": "Success"})
			return
		}

		ctx.JSON(500, gin.H{"message": "Username or password did not match"})
	}
}

func Find(usersCollection *mongo.Collection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user bson.M
		err := usersCollection.FindOne(context.TODO(), bson.M{}).Decode(&user)
		if err != nil {
			panic(err)
		}

		ctx.JSON(200, user)
	}
}
