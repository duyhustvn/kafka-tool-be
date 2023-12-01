package kafkarest

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	kafkasvc "kafkatool/internal/modules/kafka/service"

	"github.com/gorilla/mux"
)

type kafkaHandlers struct {
	router           *mux.Router
	kafkaSvc         kafkasvc.KafkaSvc
	log              logger.Logger
	cfg              *config.Config
	metricsCollector metrics.IMetricCollector
}

func NewKafkaHandlers(router *mux.Router, log logger.Logger, cfg *config.Config, kafkaSvc *kafkasvc.KafkaSvc, metricCollector metrics.IMetricCollector) *kafkaHandlers {
	return &kafkaHandlers{router: router, log: log, cfg: cfg, kafkaSvc: *kafkaSvc, metricsCollector: metricCollector}
}
