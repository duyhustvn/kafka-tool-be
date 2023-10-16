package healthcheckrest

import (
	"kafkatool/internal/common"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	healthchecksvc "kafkatool/internal/modules/healthcheck/service"
	"net/http"

	"github.com/gorilla/mux"
)

type healthcheckHandlers struct {
	router           *mux.Router
	log              logger.Logger
	cfg              *config.Config
	metricsCollector metrics.IMetricCollector
	healthCheckSvc   *healthchecksvc.HealthCheckSvc
}

func NewHealthCheckHandlers(router *mux.Router, log logger.Logger, cfg *config.Config, healthcheckSvc *healthchecksvc.HealthCheckSvc, metricCollector metrics.IMetricCollector) *healthcheckHandlers {
	return &healthcheckHandlers{router: router, log: log, cfg: cfg, healthCheckSvc: healthcheckSvc, metricsCollector: metricCollector}
}

func (handler *healthcheckHandlers) HealthCheckHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler.healthCheckSvc.HealthCheck()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		common.ResponseOk(w, http.StatusOK, nil)
	}
}
