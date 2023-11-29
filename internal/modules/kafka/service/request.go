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

func (svc KafkaSvc) SaveRequest(ctx context.Context, request kafkamodel.Request) error {
	if err := svc.sqlRepo.SaveRequest(ctx, request); err != nil {
		return err
	}
	return nil
}
