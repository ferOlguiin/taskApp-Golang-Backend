package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("El archivo .env no se encuentra")
	// }
	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatal("la uri MongoDB no puede ser un string vacio!")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de datos conectada correctamente")

	return client

}
