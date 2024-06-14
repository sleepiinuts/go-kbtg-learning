package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() error {
	password := "HDrLv7gRVMJ8BJaX"
	conn := fmt.Sprintf("mongodb+srv://teerapat:%s@cluster0.5m0mj.mongodb.net/%s?retryWrites=true&w=majority", password, "sample_airbnb")
	log.Println("conn:", conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
	if err != nil {
		fmt.Println("connect-error: ", err)
		return err
	}

	coll := client.Database("sample_airbnb").Collection("listingsAndReviews")
	cur, err := coll.Find(ctx, bson.M{}, options.Find().SetLimit(5))
	if err != nil {
		return err
	}

	var abs []AirBNB
	if err := cur.All(ctx, &abs); err != nil {
		return err
	}

	log.Println(len(abs))
	for _, a := range abs {
		log.Printf("airbnb: %+v", a)
		log.Print("\n\n\n")
	}
	return nil
}

type AirBNB struct {
	Name        string `bson:"name"`
	Space       string `bson:"space"`
	Description string `bson:"description"`
}
