package mongoclient

import (
	"context"
	"github.com/hasanbakirci/mongodb-kafka-connector/config"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDb() (db *mongo.Database, err error) {
	log.Infof("Mongo:Connection Uri:%s", config.Config.MONGOURL)
	clientOptions := options.Client().ApplyURI(config.Config.MONGOURL)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Errorf("Mongo: couldn't connect to mongo: %v", err)
		return db, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Errorf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Errorf("Mongo: Client Ping error", err)
	}
	db = client.Database(config.Config.DATABASENAME)
	return db, err

}
