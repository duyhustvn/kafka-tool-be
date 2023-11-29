package kafkasvc

import (
	"context"
	"fmt"
)

func (svc *KafkaSvc) ListTopic(ctx context.Context) ([]string, error) {
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
