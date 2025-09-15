package kafkaclient

import (
	"fmt"
	"kafkatool/internal/config"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// GetNewKafkaReader create new kafka reader
func GetNewKafkaReader(
	kafkaCfg config.Kafka,
	groupTopics []string,
	groupID string,
) *kafka.Reader {
	cfg := kafka.ReaderConfig{
		Brokers:                kafkaCfg.Brokers,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	}

	fmt.Printf("[GetNewKafkaReader] using mechanism %s username: %s", kafkaCfg.SaslMechanism, kafkaCfg.Username)
	if strings.ToUpper(kafkaCfg.SaslMechanism) == "SASL_PLAIN" {
		mechanism := plain.Mechanism{
			Username: kafkaCfg.Username,
			Password: kafkaCfg.Password,
		}
		cfg.Dialer = &kafka.Dialer{
			SASLMechanism: mechanism,
		}
	}

	return kafka.NewReader(cfg)
}
