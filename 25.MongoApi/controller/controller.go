package controller

import (
	"context"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// connection string with mongo Db
const connectionString = "mongodb+srv://sonu517825:m0ww1dng9uqrz0ge@cluster0.wgtiy.mongodb.net/group13Database?retryWrites=true&w=majority"

const dbName = "NetFlix"
const collectionName = "watchlist"

// most importamt pointer to mongo collection
var collection *mongo.Collection

//connect with mongoDb

func init() {
	// special method in go lang which runs only for the first time
	// onlyone time

	//client connection

	clinetOptions := options.Client()

	client, err := mongo.Connect(context.TODO(), clinetOptions)

	if err != nil {

	}

	fmt.Println(" mongo db Connected ")

	// creating db and collection

	collection = client.Database(collectionName).Collection(dbName)

	//collection instance

	fmt.Println((" collection instance is ready "))

}

// mongo helpers
// usually go into another folder

func AddOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		//panic(err) or
		log.Fatal(err) // more controlled version
	}

	fmt.Println("movie succesfully  inserted with id ", inserted.InsertedID)

}

func UpdateWatchedOrNot(movieId string) {

	// thing is that we need top get the id from string and convert it into objectId
	//the following syntax is provided by mongo db

	id, _ := primitive.ObjectIDFromHex(movieId)

	//everything in mongo DB bson (json+some extra data )

	filter := bson.M{"_id", id}                       // somethig findiing id in mongo db
	update := bson.M{"$set": bson.M{"watched": true}} //also can used bsonM

	result, err := collection.UpdateOne(context.Background(), filter, update)

	//something like in mongoose
	//movieModel.findOneAndUpdate({"id",_id},{"$set":{"watched":true}})

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("modifie count", result.ModifiedCount)

}

func DeleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id", id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter) //this is not an ORM
	// this is raw mongo db command
	// it returns a delete count   that how mny values are actually being deleted
	// also it directly returns a integer
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("movie got deleted with delete count ", deleteCount)

}

//delete all record from mongodB

func DeleteAllMovie() {
	filter := bson.D{{}} // passing inside nothing that means delete all
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" numbers of movie deleted ", deleteResult.DeletedCount)
}

//get all movies

func GetAllmovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{{}})

	// it will return an object with lots of values so we need to iterate over that
	// and make  it fruitfull

	if err != nil {
		panic(err)
	}

	var movies []primitive.M

	// like a while loop till the cur(object ) having something on ot
	for cur.Next(context.Background()) {

		var movie bson.M // isme movies aaengi
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}

	defer cur.Close(context.Background())

	return movies

}

func ServeHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> welcole to the home route </h1>"))

}
