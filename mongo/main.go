package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var personCollection *mongo.Collection
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	personCollection = client.Database("test_db").Collection("person")
}

type Person struct {
	Name string
	Age  int
	City string
}

func main() {

	AddOneperson()
	// AddManyPerson()
	// FindAllPerson()
	// FindPodcastWithFilter()
	// FindWithOptions()
	// UpdateDocument()
}

func AddOneperson() {

	person1 := Person{"apple", 34, "Cape Town"}

	insertResult, err := personCollection.InsertOne(ctx, person1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted with ID : ", insertResult.InsertedID)
}

func AddManyPerson() {

	person2 := Person{"ball", 32, "Nairobi"}
	person3 := Person{"cat", 31, "Nairobi"}

	temp1 := []interface{}{person2, person3}

	insertManyResult, err := personCollection.InsertMany(ctx, temp1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

//getting all the documents in batches, recomended when there are lot of data in database change D to M to get as maps
func FindAllPerson() {
	cursor, err := personCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var person []bson.M
	if err = cursor.All(ctx, &person); err != nil {
		log.Fatal(err)
	}
	for key, val := range person {
		fmt.Println(key, val)
	}
}

/* query documents with filter */
func FindPodcastWithFilter() {
	filterCursor, err := personCollection.Find(ctx, bson.M{"name": "apple"})
	if err != nil {
		log.Fatal(err)
	}
	var podcast []bson.M
	if err = filterCursor.All(ctx, &podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)
}

//find with options
func FindWithOptions() {
	findOptions := options.Find()                   // build a `findOptions`
	findOptions.SetSort(map[string]int{"when": -1}) // reverse order by `when`
	findOptions.SetSkip(0)                          // skip whatever you want, like `offset` clause in mysql
	findOptions.SetLimit(2)                         // like `limit` clause in mysql

	filterCursor, err := personCollection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	var podcast []bson.M
	if err = filterCursor.All(ctx, &podcast); err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcast)
}

func UpdateDocument() {

	filter := bson.D{{"name", "apple"}}

	update := bson.D{{"$set",
		bson.D{
			{"city", "arsikere"},
		},
	}}

	updateResult, err := personCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

}
