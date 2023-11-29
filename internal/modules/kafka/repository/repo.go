package kafkarepo

import (
	"context"
	kafkamodel "kafkatool/internal/modules/kafka/models"
)

type ISqlRepo interface {
	ListRequest(ctx context.Context) ([]kafkamodel.Request, error)
	SaveRequest(ctx context.Context, request kafkamodel.Request) error
}
