package kafkareqsvc

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	kafkareqrepo "kafkatool/internal/modules/kafka_request/repository"
)

type KafkaSvc struct {
	sqlRepo kafkareqrepo.ISqlRepo
	log     logger.Logger
	cfg     config.Config
}

func NewKafkaRequestSvc(sqlRepo kafkareqrepo.ISqlRepo, log logger.Logger, cfg config.Config) *KafkaSvc {
	return &KafkaSvc{sqlRepo: sqlRepo, log: log, cfg: cfg}
}
