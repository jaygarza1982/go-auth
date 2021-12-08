package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jaygarza1982/go-auth/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Getenv("MONGO_URL")
	dbName := os.Getenv("MONGO_INITDB_DATABASE")

	fmt.Printf("Attemping to connect to MongoDB on %v\n", uri)
	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("Could not connect to MongoDB! %v", err)
			panic(err)
		}
	}()

	db := client.Database(dbName)
	usersCollection := db.Collection("users")

	ginServer := gin.Default()

	sessionSecret := os.Getenv("SESSION_SECRET")
	store := cookie.NewStore([]byte(sessionSecret))
	ginServer.Use(sessions.Sessions("session", store))

	ginServer.GET("/", func(c *gin.Context) {
		status := "Status is OK"
		c.JSON(200, gin.H{"message": status})
	})

	ginServer.POST("/auth/find", controllers.Find(usersCollection))
	ginServer.POST("/auth/register", controllers.Register(usersCollection))
	ginServer.POST("/auth/login", controllers.Login(usersCollection))
	ginServer.POST("/auth/logout", controllers.Logout())
	ginServer.GET("/auth/check", controllers.Check())

	ginServer.Run(":8080")
}
