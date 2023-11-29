package kafkasvc

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	kafkarepo "kafkatool/internal/modules/kafka/repository"
	kafkaclient "kafkatool/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

type KafkaSvc struct {
	kafkaConn          *kafka.Conn
	kafkaProducer      kafkaclient.Producer
	kafkaConsumerGroup kafkaclient.ConsumerGroup
	sqlRepo            kafkarepo.ISqlRepo
	log                logger.Logger
	cfg                config.Config
}

func NewKafkaSvc(kafkaConn *kafka.Conn, kafkaProducer kafkaclient.Producer, kafkaConsumerGroup kafkaclient.ConsumerGroup, sqlRepo kafkarepo.ISqlRepo, log logger.Logger, cfg config.Config) *KafkaSvc {
	return &KafkaSvc{kafkaConn: kafkaConn, kafkaProducer: kafkaProducer, kafkaConsumerGroup: kafkaConsumerGroup, sqlRepo: sqlRepo, log: log, cfg: cfg}
}
