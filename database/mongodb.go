package database

import (
	"context"
	"fmt"
	"log"

	"github.com/prabhaw/golang-startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var Collection *mongo.Collection
var Client *mongo.Client

func Connect() error {
	var err error
var connectionString string

// Load Credenials
dbHost := config.Config("DB_HOST")
dbPort := config.Config("DB_PORT")
dbName := config.Config("DB_NAME")
collName := config.Config("COLLECTION_NAME")
dbUser := config.Config("DB_USER")
dbPass := config.Config("DB_PASS")
isSRV := config.Config("SRV")


if(isSRV == "true"){
	connectionString = fmt.Sprintf("mongodb+srv://%s:%s@%s",dbUser, dbPass,dbHost)
} else{
	connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost,dbPort)
}


// Create a new Client
Client, err = mongo.Connect(context.TODO(),options.Client().ApplyURI(connectionString))
if err != nil{
	log.Fatalln("Connect:", err)
	return err
}

Collection = Client.Database(dbName).Collection(collName)
log.Println("Open database connection and loaded collection")
return nil
}