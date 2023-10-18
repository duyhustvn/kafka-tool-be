package kafkaclient

import (
	"context"
	"kafkatool/internal/logger"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	Brokers []string
	log     logger.Logger
	w       *kafka.Writer
}

func NewProducer(brokers []string, log logger.Logger) *producer {
	return &producer{Brokers: brokers, log: log, w: GetNewKafkaWriter(brokers)}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
