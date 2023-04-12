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

func deleteOneMovie(movieId string)  {
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"$_id":id}
	deleteCount, err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Movie to delete on delete count",deleteCount)
}

//delete all records from mongoDb

func deleteAllMovie() int64  {
	filter:=bson.D{{}}
	deleteResult, err:= collection.DeleteMany(context.Background(),filter,nil)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies to be delete",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cur,err:= collection.Find(context.Background(),bson.D{{}})
if err!=nil {
	log.Fatal(err)
}
defer cur.Close(context.Background())
var movies []primitive.M

for cur.Next(context.Background()){
	var movie bson.M
	err:=cur.Decode(&movie)
	if err!=nil {
		log.Fatal(err)
	}
	movies = append(movies, movie)
}
return movies //It uses the Find method on a MongoDB collection to retrieve all documents (bson.D{{}} represents an empty filter, which matches all documents).
//It checks for any errors that occurred during the Find operation, and if there was an error, it logs it and exits the function with a fatal error.
//It defers the closing of the MongoDB cursor (defer cur.Close(context.Background())) to ensure that the cursor is closed when the function finishes execution.
//It creates an empty slice ([]primitive.M) to hold the retrieved documents.
//It loops through the cursor with cur.Next(context.Background()), which returns true if there is another document to be read.
//For each document, it decodes it into a bson.M struct.
//If there was an error during decoding, it logs it and exits the function with a fatal error.
//It appends the decoded document to the movies slice.
//Once all documents have been retrieved and appended to the movies slice, it returns the slice.

}
