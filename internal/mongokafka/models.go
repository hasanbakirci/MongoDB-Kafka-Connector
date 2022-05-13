package mongokafka

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type KafkaLog struct {
	ID        primitive.ObjectID `bson:"_id"`
	LogId     string             `bson:"logId"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func Create(k KafkaLog) (*KafkaLog, error) {
	return &KafkaLog{
		ID:        GenerateObjectId(),
		LogId:     k.LogId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func GenerateObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}

type CreateKafkaLogRequest struct {
	LogId string `json:"log_id"`
}

func (receiver *CreateKafkaLogRequest) ToKafkaLog() *KafkaLog {
	return &KafkaLog{
		LogId: receiver.LogId,
	}
}
