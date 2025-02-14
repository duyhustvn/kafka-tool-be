package kafkareqsvc

import (
	"context"
	kafkareqmodel "kafkatool/internal/modules/kafka_request/models"
)

func (svc KafkaSvc) ListRequest(ctx context.Context) ([]kafkareqmodel.Request, error) {
	requests, err := svc.sqlRepo.ListRequest(ctx)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (svc KafkaSvc) CreateRequest(ctx context.Context, request kafkareqmodel.Request) (int64 /*insertedId*/, error) {
	insertedId, err := svc.sqlRepo.CreateRequest(ctx, request)
	if err != nil {
		return -1, err
	}
	return insertedId, nil
}

func (svc KafkaSvc) UpdateRequest(ctx context.Context, requestID int, request kafkareqmodel.Request) error {
	if err := svc.sqlRepo.UpdateRequest(ctx, requestID, request); err != nil {
		return err
	}
	return nil
}

func (svc KafkaSvc) DeleteRequest(ctx context.Context, requestID int) error {
	if err := svc.sqlRepo.DeleteRequest(ctx, requestID); err != nil {
		return err
	}
	return nil
}
