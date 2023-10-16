package healthchecksvc

import "kafkatool/internal/logger"

type HealthCheckSvc struct {
	log logger.Logger
}

func NewHealthCheckSvc(log logger.Logger) (*HealthCheckSvc, error) {
	return &HealthCheckSvc{log: log}, nil
}

func (svc *HealthCheckSvc) HealthCheck() error {
	return nil
}
