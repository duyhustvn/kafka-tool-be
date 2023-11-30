package kafkarequestres

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	kafkareqsvc "kafkatool/internal/modules/kafka_request/service"

	"github.com/gorilla/mux"
)

type kafkaHandlers struct {
	router           *mux.Router
	kafkaSvc         kafkareqsvc.KafkaSvc
	log              logger.Logger
	cfg              config.Config
	metricsCollector metrics.IMetricCollector
}

func NewKafkaRequestHandlers(router *mux.Router, log logger.Logger, cfg config.Config, kafkaSvc *kafkareqsvc.KafkaSvc, metricCollector metrics.IMetricCollector) *kafkaHandlers {
	return &kafkaHandlers{router: router, log: log, cfg: cfg, kafkaSvc: *kafkaSvc, metricsCollector: metricCollector}
}
