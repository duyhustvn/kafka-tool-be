package kafkarepo

import (
	"context"
	kafkamodel "kafkatool/internal/modules/kafka/models"
)

type ISqlRepo interface {
	ListRequest(ctx context.Context) ([]kafkamodel.Request, error)
	CreateRequest(ctx context.Context, request kafkamodel.Request) error
	UpdateRequest(ctx context.Context, requestID int, request kafkamodel.Request) error
}
