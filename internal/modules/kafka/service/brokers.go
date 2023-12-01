package kafkasvc

import (
	"context"
	"fmt"
	"kafkatool/internal/config"
	kafkaclient "kafkatool/pkg/kafka"
	"strings"
)

func (svc *KafkaSvc) ConnectKafkaBrokers(ctx context.Context, brokers string, cfg *config.Kafka) error {
	newBrokers := strings.Split(brokers, ",")
	cfg.Brokers = newBrokers

	// close old kafka connection
	if svc.kafkaConn != nil {
		svc.kafkaConn.Close()
	}

	kafkaConn, err := kafkaclient.NewKafkaConnection(ctx, cfg)
	if err != nil {
		// reset kafka connection

		svc.kafkaConn = nil
		svc.kafkaProducer = nil
		svc.kafkaConsumerGroup = nil

		return fmt.Errorf("failed to connect to kafka brokers %+v", err)
	}

	svc.kafkaConn = kafkaConn
	svc.kafkaProducer = kafkaclient.NewProducer(cfg.Brokers, svc.log)
	svc.kafkaConsumerGroup = kafkaclient.NewConsumerGroup(cfg.Brokers, cfg.GroupID, svc.log)

	return nil
}

func (svc *KafkaSvc) IsConnectedToKafkaBrokers(ctx context.Context) bool {
	return svc.kafkaConn != nil
}
