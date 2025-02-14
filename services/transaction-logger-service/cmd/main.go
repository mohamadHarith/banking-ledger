package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	uri := fmt.Sprintf("")
	mongo.Connect(options.Client().ApplyURI(uri))
}
