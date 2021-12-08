package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(usersCollection *mongo.Collection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		_, err := usersCollection.InsertOne(context.TODO(), bson.M{"username": "username123"})
		if err != nil {
			ctx.JSON(200, gin.H{})
		}
		if err == nil {
			panic(err)
		}
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

func Login(usersCollection *mongo.Collection) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	}
}
