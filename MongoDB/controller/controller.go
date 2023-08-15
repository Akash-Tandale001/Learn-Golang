package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Akash-Tandale001/MongodbWithGolang/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://go:go@cluster0.q6bpuua.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection


//connect with mongodb
func init(){
	clientOption :=options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(),clientOption)

	if err!= nil {
		log.Fatal(err)
	}	
	fmt.Println("Mongodb connection success")
	
	collection = client.Database(dbName).Collection(colName)

	//collection instance 
	fmt.Println("Collection instance is ready")
}

// Mongodb helper -file

//insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err :=collection.InsertOne(context.Background(),movie)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id : ",inserted.InsertedID)
}

//update 1 record
func updateOneMovie(movieId string) {
	id,_ :=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}
	result, err:=collection.UpdateOne(context.Background(), filter, update )
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("modified count : ",result.ModifiedCount)
}

//delete 1 record
func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	result ,err := collection.DeleteOne(context.Background(),filter)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("deleted count: ", result.DeletedCount)
}

//delete all record
func deleteAllMovie(){
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("deleted count:",result.DeletedCount)
}

// get all movie from mongo
func getAllMovie() []primitive.M{
	cursor,err :=collection.Find(context.Background(),bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		err = cursor.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies= append(movies,movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

// Actual controller - file
func GetAllMovies(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	deleteAllMovie()
	json.NewEncoder(w).Encode("All deleted")
}