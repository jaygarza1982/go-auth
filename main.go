package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	uri := os.Getenv("MONGO_URL")
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
			fmt.Println("Could not connect to MongoDB!")
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")

	type test struct {
		Title string
	}

	s := test{Title: "Back to the Future"}
	if result, err := coll.InsertOne(context.TODO(), s); err == nil {
		fmt.Printf("Result of insert was %v", result)
	}
	if err != nil {
		panic(err)
	}

	ginServer := gin.Default()

	ginServer.GET("/", func(c *gin.Context) {
		status := "Status is OK"
		c.JSON(200, gin.H{"message": status})
	})

	ginServer.Run(":8080")
}
