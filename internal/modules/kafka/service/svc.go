package kafkasvc

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	kafkaclient "kafkatool/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

type KafkaSvc struct {
	kafkaConn          *kafka.Conn
	kafkaProducer      kafkaclient.Producer
	kafkaConsumerGroup kafkaclient.ConsumerGroup
	log                logger.Logger
	cfg                config.Config
}

func NewKafkaSvc(kafkaConn *kafka.Conn, kafkaProducer kafkaclient.Producer, kafkaConsumerGroup kafkaclient.ConsumerGroup, log logger.Logger, cfg config.Config) *KafkaSvc {
	return &KafkaSvc{kafkaConn: kafkaConn, kafkaProducer: kafkaProducer, kafkaConsumerGroup: kafkaConsumerGroup, log: log, cfg: cfg}
}
