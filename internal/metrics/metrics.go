package metrics

import (
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics/adapter"
	"kafkatool/internal/metrics/adapter/noop"
	"kafkatool/internal/metrics/adapter/statsd"
	"time"
)

// IMetricCollector
type IMetricCollector interface {
	RegisterCounter(adapter.CollectorOptions) adapter.Counter
	RegisterGauge(adapter.CollectorOptions) adapter.Gauge
	RegisterTimer(adapter.CollectorOptions) adapter.Timer
	Shutdown()
}

func NewCollector(cfg *config.Config, log logger.Logger) IMetricCollector {
	monitoringConfig := &cfg.Monitoring
	isStatsd := monitoringConfig != nil && monitoringConfig.Statsd != nil

	if isStatsd {
		flushDuration := 100 * time.Millisecond
		if monitoringConfig.Statsd.FlushPeriod > 0 {
			flushDuration = time.Duration(monitoringConfig.Statsd.FlushPeriod) * time.Millisecond
		}

		return statsd.NewCollector(monitoringConfig.Statsd.Addr, monitoringConfig.Statsd.Prefix, log, flushDuration)
	}

	log.Info("config_skipping_empty_metrics_provider")
	return noop.NewCollector()
}
