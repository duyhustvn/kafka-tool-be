package kafkasvc

import (
	"context"
	kafkamodel "kafkatool/internal/modules/kafka/models"
)

func (svc *KafkaSvc) ListRequest(ctx context.Context) ([]kafkamodel.Request, error) {
	return []kafkamodel.Request{
		{Title: "Request 1", Topic: "sample_topic", Quantity: 11, Type: "json", Message: `{"key1": "value1"}`},
		{Title: "Request 2", Topic: "sample_topic", Quantity: 12, Type: "json", Message: `{"key2": "value2"}`},
	}, nil
}
