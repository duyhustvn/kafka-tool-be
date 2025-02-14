package kafkareqrepo

import (
	"context"
	kafkareqmodel "kafkatool/internal/modules/kafka_request/models"
)

type ISqlRepo interface {
	ListRequest(ctx context.Context) ([]kafkareqmodel.Request, error)
	CreateRequest(ctx context.Context, request kafkareqmodel.Request) (int64 /*insertedId*/, error)
	UpdateRequest(ctx context.Context, requestID int, request kafkareqmodel.Request) error
	DeleteRequest(ctx context.Context, requestID int) error
}
