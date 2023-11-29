package kafkarepo

import (
	"context"
	kafkamodel "kafkatool/internal/modules/kafka/models"
)

type ISqlRepo interface {
	SaveRequest(ctx context.Context, request kafkamodel.Request) error
}
