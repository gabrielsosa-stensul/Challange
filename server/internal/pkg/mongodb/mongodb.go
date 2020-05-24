package mongodb

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	environmentloader "github.com/MarianoArias/Challange/server/internal/pkg/environment-loader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// init establishes the connection to mongodb with the credentials defined in the 
// environment variables and set up a mongo client.
func init() {
	environmentloader.Load()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+""))

	err := conn.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("\033[97;41m%s\033[0m\n", "### MongoDB connection error: "+err.Error()+" ###")
	} else {
		client = conn
		log.Printf("\033[97;42m%s\033[0m\n", "### MongoDB connection established ###")
	}
}

// Ping returns an error if the connection is down.
func Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Ping(ctx, nil)

	if err != nil {
		return errors.New("Could not ping")
	}

	return nil
}

// GetClient returns a mongo client
func GetClient() *mongo.Client {
	return client
}

// GetDatabase returns a mongo database that is defined in the environment 
// variables.
func GetDatabase() *mongo.Database {
	return client.Database(os.Getenv("DATABASE_NAME"))
}
