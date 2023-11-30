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

func (svc KafkaSvc) CreateRequest(ctx context.Context, request kafkareqmodel.Request) error {
	if err := svc.sqlRepo.CreateRequest(ctx, request); err != nil {
		return err
	}
	return nil
}

func (svc KafkaSvc) UpdateRequest(ctx context.Context, requestID int, request kafkareqmodel.Request) error {
	if err := svc.sqlRepo.UpdateRequest(ctx, requestID, request); err != nil {
		return err
	}
	return nil
}
