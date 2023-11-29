package kafkasvc

import (
	"context"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	kafkamodel "kafkatool/internal/modules/kafka/models"
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

func (svc KafkaSvc) SaveRequest(ctx context.Context, request kafkamodel.Request) error {
	if err := svc.sqlRepo.SaveRequest(ctx, request); err != nil {
		return err
	}
	return nil
}
