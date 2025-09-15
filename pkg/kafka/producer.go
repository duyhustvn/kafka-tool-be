package kafkaclient

import (
	"context"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	w *kafka.Writer
}

func NewProducer(
	kafkaCfg config.Kafka,
	log logger.Logger,
) *producer {
	return &producer{
		w: GetNewKafkaWriter(kafkaCfg),
	}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
