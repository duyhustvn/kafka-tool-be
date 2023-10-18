package config

import (
	"strconv"
	"strings"
)

type Kafka struct {
	Brokers  []string
	GroupID  string
	PoolSize int
}

func (k *Kafka) GetKafkaEnv() *Kafka {
	brokers := GetEnv("KAFKA_BROKERS")
	k.Brokers = strings.Split(brokers, ",")
	k.GroupID = GetEnv("KAFKA_GROUP_ID")
	ps, err := strconv.Atoi(GetEnv("KAFKA_POOL_SIZE"))
	if err != nil {
		k.PoolSize = 30
	} else {
		k.PoolSize = ps
	}

	return k
}
