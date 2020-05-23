package mongodb

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	environmentloader "github.com/MarianoArias/challange-api/internal/pkg/environment-loader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

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

func Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Ping(ctx, nil)

	if err != nil {
		return errors.New("Could not ping")
	}

	return nil
}

func GetClient() *mongo.Client {
	return client
}

func GetDatabase() *mongo.Database {
	return client.Database(os.Getenv("DATABASE_NAME"))
}
