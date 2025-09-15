package kafkaclient

import (
	"fmt"
	"kafkatool/internal/config"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// GetNewKafkaWriter create new kafka producer
func GetNewKafkaWriter(
	kafkaCfg config.Kafka,
) *kafka.Writer {
	cfg := kafka.WriterConfig{
		Brokers:      kafkaCfg.Brokers,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}

	fmt.Printf("[GetNewKafkaWriter] using mechanism %s username: %s \n", kafkaCfg.SaslMechanism, kafkaCfg.Username)

	if strings.ToUpper(kafkaCfg.SaslMechanism) == "SASL_PLAIN" {
		mechanism := plain.Mechanism{
			Username: kafkaCfg.Username,
			Password: kafkaCfg.Password,
		}

		cfg.Dialer = &kafka.Dialer{
			SASLMechanism: mechanism,
		}
	}

	w := kafka.NewWriter(cfg)

	return w
}
