package kafkasvc

import (
	"context"
	kafkaclient "kafkatool/pkg/kafka"
)

func (svc *KafkaSvc) ConsumeMessage(ctx context.Context, topics []string, poolSize int, worker kafkaclient.Worker) {
	go svc.kafkaConsumerGroup.ConsumeTopic(ctx, topics, poolSize, worker)
}
