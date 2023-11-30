package kafkasvc

import (
	"context"
	kafkamodel "kafkatool/internal/modules/kafka/models"
)

func (svc KafkaSvc) ListRequest(ctx context.Context) ([]kafkamodel.Request, error) {
	requests, err := svc.sqlRepo.ListRequest(ctx)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (svc KafkaSvc) CreateRequest(ctx context.Context, request kafkamodel.Request) error {
	if err := svc.sqlRepo.CreateRequest(ctx, request); err != nil {
		return err
	}
	return nil
}

func (svc KafkaSvc) UpdateRequest(ctx context.Context, requestID int, request kafkamodel.Request) error {
	if err := svc.sqlRepo.UpdateRequest(ctx, requestID, request); err != nil {
		return err
	}
	return nil
}
