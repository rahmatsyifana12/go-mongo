package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name	string	`bson:"name" json:"name"`
	Age		int		`bson:"age" json:"age"`
	Address	string	`bson:"address,omitempty" json:"address,omitempty"`
	Phone	string	`bson:"phone" json:"phone"`
}

type Item struct {
	ItemID          int     `bson:"item_id" json:"item_id"`
	Description     string  `bson:"description" json:"description"`
	Price           float32 `bson:"price" json:"price"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("sample_mflix")
	user_collection := db.Collection("users")

	new_users := []interface{}{
		User{
			Name: "John Doe",
			Age: 20,
			Address: "Kampung gajah",
		},
		User{
			Name: "Mary Jane",
			Age: 21,
			Phone: "08123456789",
		},
	}

	usr_res, err := user_collection.InsertMany(ctx, new_users)
	if err != nil {
		panic(err)
	}
	fmt.Println(usr_res)

	// var result bson.M
	// err = user_collection.FindOne(context.TODO(), bson.D{{"name", "rahmat"}}).Decode(&result)
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// user_res, _ := user_collection.InsertOne(ctx, bson.D{
	// 	{Key: "name", Value: "Rahmat"},
	// })
	// fmt.Println(user_res)

	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)
}
