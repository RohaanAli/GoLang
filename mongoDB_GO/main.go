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

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://*******:*******@rcluster-5hsrk.mongodb.net/products?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	Shopdb := client.Database("Shopdb")

	products := Shopdb.Collection("products")

	//Reading ALL
	// cursor, err := products.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pro := []bson.M{}

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = cursor.All(ctx, &pro); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, p := range pro {
	// 	fmt.Println(p["name"])
	// }

	//Reading in batches for bigger db and with filter
	// cursor, err := products.Find(ctx, bson.M{"name": "Pencil"})

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(ctx)

	// for cursor.Next(ctx) {
	// 	var p bson.M
	// 	if err = cursor.Decode(&p); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(p)
	// }

	//FindOne...........gets the first one

	var p bson.M
	if err = products.FindOne(ctx, bson.M{}).Decode(&p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

}
