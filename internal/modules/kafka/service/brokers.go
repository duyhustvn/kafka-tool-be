package kafkasvc

import (
	"context"
	"fmt"
	"kafkatool/internal/config"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	kafkaclient "kafkatool/pkg/kafka"
	"strings"
)

func (svc *KafkaSvc) ConnectKafkaBrokers(
	ctx context.Context,
	body kafkamodel.BrokersConfig,
) error {
	brokersUrl := body.Url
	brokers := strings.Split(brokersUrl, ",")

	kafkaCfg := config.Kafka{
		Brokers: brokers,
	}

	if strings.TrimSpace(body.Username) != "" && strings.TrimSpace(body.Password) != "" {
		kafkaCfg.Username = body.Username
		kafkaCfg.Password = body.Password
		kafkaCfg.SaslMechanism = "SASL_PLAIN"
	}

	// close old kafka connection
	if svc.kafkaConn != nil {
		svc.kafkaConn.Close()
	}

	svc.log.Infof("Trying to connect to kafka brokers: %+v", brokers)
	kafkaConn, err := kafkaclient.NewKafkaConnection(ctx, &kafkaCfg)
	if err != nil {
		// reset kafka connection

		svc.kafkaConn = nil
		svc.kafkaProducer = nil
		svc.kafkaConsumerGroup = nil

		return fmt.Errorf("failed to connect to kafka brokers %+v", err)
	}

	svc.kafkaConn = kafkaConn
	svc.kafkaProducer = kafkaclient.NewProducer(kafkaCfg, svc.log)
	svc.kafkaConsumerGroup = kafkaclient.NewConsumerGroup(kafkaCfg, "kafka-tool", svc.log)
	svc.log.Infof("Connected to kafka brokers: %+v success", brokers)

	return nil
}

func (svc *KafkaSvc) IsConnectedToKafkaBrokers(ctx context.Context) bool {
	return svc.kafkaConn != nil
}
