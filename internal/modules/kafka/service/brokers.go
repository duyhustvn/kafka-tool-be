package kafkasvc

import (
	"context"
	"fmt"
	"kafkatool/internal/config"
	kafkaclient "kafkatool/pkg/kafka"
)

func (svc *KafkaSvc) ConnectKafkaBrokers(ctx context.Context, brokers []string, cfg *config.Kafka) error {
	cfg.Brokers = brokers

	// close old kafka connection
	if svc.kafkaConn != nil {
		svc.kafkaConn.Close()
	}

	svc.log.Infof("Trying to connect to kafka brokers: %+v", brokers)
	kafkaConn, err := kafkaclient.NewKafkaConnection(ctx, cfg)
	if err != nil {
		// reset kafka connection

		svc.kafkaConn = nil
		svc.kafkaProducer = nil
		svc.kafkaConsumerGroup = nil

		return fmt.Errorf("failed to connect to kafka brokers %+v", err)
	}

	svc.log.Infof("Connected to kafka brokers: %+v success", brokers)
	svc.kafkaConn = kafkaConn
	svc.kafkaProducer = kafkaclient.NewProducer(cfg.Brokers, svc.log)
	svc.kafkaConsumerGroup = kafkaclient.NewConsumerGroup(cfg.Brokers, cfg.GroupID, svc.log)

	return nil
}

func (svc *KafkaSvc) IsConnectedToKafkaBrokers(ctx context.Context) bool {
	return svc.kafkaConn != nil
}
