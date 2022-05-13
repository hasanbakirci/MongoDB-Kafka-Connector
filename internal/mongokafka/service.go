package mongokafka

import (
	"context"
	"github.com/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, request CreateKafkaLogRequest) (string, error)
}

type service struct {
	repository Repository
}

func (s service) Create(ctx context.Context, request CreateKafkaLogRequest) (string, error) {
	result, err := s.repository.Create(ctx, request.ToKafkaLog())
	if err != nil {
		return "", errors.Wrap(err, "Service: Failed to create order")
	}
	return result.String(), nil
}

func NewService(r Repository) Service {
	return &service{repository: r}
}
