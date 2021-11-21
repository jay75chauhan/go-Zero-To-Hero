package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jay75chauhan/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://go:@cluster0.bhgvg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//most important

   var collection *mongo.Collection


//conect with mongodb collecti


func init() {
	//clint option

	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb

   client,err := mongo.Connect(context.TODO(), clientOptions)


   if err != nil {
	   log.Fatal(err)
   }
   
   fmt.Println("MongoDB connection succes")
 
    collection = client.Database(dbName).Collection(colName)


	//collectiom instance
	fmt.Println("collection instant redy")


}

// MONGODB helpers - file 


// insert 1 record


func insertOneMovie(movie model.Netflix) {

 inserted, err :=   collection.InsertOne(context.Background(),movie)

  if err != nil {
	  log.Fatal(err)
  }

  fmt.Println("Inserted 1 movie successfully",inserted.InsertedID)

}

// update 1 record 

func updateOneMovie(movieId string){

  id,_:=	primitive.ObjectIDFromHex(movieId)

  filter :=  bson.M{"_id":id}

  update := bson.M{"$set": bson.M{"watched": true}}

  result,err :=collection.UpdateOne(context.Background(),filter,update)


  if err != nil {
	  log.Fatal(err)
  }

  fmt.Println("modified count:",result.ModifiedCount)

}

// delete 1 record

func deleteOneMovie(movieId string){

  id,_:=	primitive.ObjectIDFromHex(movieId)

  filter :=  bson.M{"_id":id}


  deleteCount,err := collection.DeleteOne(context.Background(),filter)


  if err != nil {
	  log.Fatal(err)
  }

  fmt.Println("Movie got deleted count:",deleteCount)

}

// delete all record

func deleteAllMovie() int64{
      
 deleteAll,err :=	collection.DeleteMany(context.Background(),bson.D{{}})

 if err != nil {
	  log.Fatal(err)
  }

  fmt.Println("number of movies dleted",deleteAll.DeletedCount)

  return deleteAll.DeletedCount

}


//get all movies from the database

func getAllMovie() []primitive.M{
 

cur,err :=	collection.Find(context.Background(),bson.D{{}})
   
  if err != nil {
	   log.Fatal(err)
  }
 
       var movies []primitive.M
      for cur.Next(context.Background()){
		  var movie bson.M
		  err := cur.Decode(&movie)
		  if err != nil {
			  log.Fatal(err)
		  }
		  movies = append(movies,movie)
	  }

	  defer cur.Close(context.Background())

	  return movies

}

// Actual controller - file



func GetMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	
	allMovies := getAllMovie()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ =json.NewDecoder(r.Body).Decode(&movie)
    insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


func MarkAsWatched(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

    params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}


func DeleteAMovie(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

    params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteALLMovie(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/x-www-from-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

     count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
