package kafkaclient

import (
	"context"
	"kafkatool/internal/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConnection(ctx context.Context, cfg *config.Config) (*kafka.Conn, error) {
	brokers := cfg.Kafka.Brokers
	return kafka.DialContext(ctx, "tcp", brokers[0])
}
