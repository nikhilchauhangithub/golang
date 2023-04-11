package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/nikhilchauhangithub/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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


//Mongodb helpers

//insert one movie

func insertOneMovie(movie model.Netflix)  {
	inserted, err := collection.InsertOne(context.Background(),movie)

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id:", inserted.InsertedID)
}

//update one record

func updateOneMovie(movieId string)  {
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:= bson.M{"$set":bson.M{"watched":true}}

	result,err:= collection.UpdateOne(context.Background(),filter,update)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ",result.ModifiedCount) //The result.ModifiedCount field contains the number of documents that were successfully modified by the update operation. In this case, since we are updating only one document, the ModifiedCount field should have a value of either 0 or 1.
}