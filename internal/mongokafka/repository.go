package mongokafka

import (
	"context"
	"github.com/hasanbakirci/mongodb-kafka-connector/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, kafkalog *KafkaLog) (primitive.ObjectID, error)
}

type kafkalogRepository struct {
	collection *mongo.Collection
}

func (k kafkalogRepository) Create(ctx context.Context, kafkalog *KafkaLog) (primitive.ObjectID, error) {
	var result, e = Create(*kafkalog)
	if e != nil {
		return primitive.ObjectID{}, e
	}
	_, err := k.collection.InsertOne(ctx, result)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.ID, err
}

func NewKafkalogRepository(db *mongo.Database) Repository {
	col := db.Collection(config.Config.DATABASENAME)
	return &kafkalogRepository{collection: col}
}
