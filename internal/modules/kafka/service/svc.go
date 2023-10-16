package kafkasvc

import (
	"fmt"
	"kafkatool/internal/logger"

	"github.com/segmentio/kafka-go"
)

type KafkaSvc struct {
	kafkaConn *kafka.Conn
	log       logger.Logger
}

func NewKafkaSvc(kafkaConn *kafka.Conn, log logger.Logger) *KafkaSvc {
	return &KafkaSvc{kafkaConn: kafkaConn, log: log}
}

func (svc *KafkaSvc) ListTopic() ([]string, error) {
	partitions, err := svc.kafkaConn.ReadPartitions()
	if err != nil {
		return nil, fmt.Errorf("cannot read partitions %+v", err)
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	topics := []string{}
	for k := range m {
		topics = append(topics, k)
	}
	return topics, nil
}
