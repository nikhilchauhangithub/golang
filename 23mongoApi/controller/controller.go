package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://nikhil:nikhil@cluster0.ajb91ha.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// most important
var collection *mongo.Collection  //reference of database


//connect with mongoDB
func init()  {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err :=mongo.Connect(context.TODO(),clientOption)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection successful")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("connection instance is ready")
}