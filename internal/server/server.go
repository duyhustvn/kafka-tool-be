package server

import (
	"fmt"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	healthchecksvc "kafkatool/internal/modules/healthcheck/service"
	healthcheckrest "kafkatool/internal/modules/healthcheck/transport/rest"
	kafkasvc "kafkatool/internal/modules/kafka/service"
	kafkarest "kafkatool/internal/modules/kafka/transport/res"
	kafkareqrepo "kafkatool/internal/modules/kafka_request/repository"
	kafkareqsvc "kafkatool/internal/modules/kafka_request/service"
	kafkarequestres "kafkatool/internal/modules/kafka_request/transport/res"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"

	sqlclient "kafkatool/pkg/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	router           *mux.Router
	Cfg              config.Config
	log              logger.Logger
	metricsCollector metrics.IMetricCollector
}

// GetApp returns main app
func GetApp() *Server {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %+v\n", err)
	}

	if err := loadVars(cfg); err != nil {
		log.Fatalf("Error loading var: %+v\n", err)
	}

	log, err := logger.GetLogger(cfg)
	if err != nil {
		log.Fatalf("Error initialize custom logger: %s\n", err)
	}

	log.Debugf("Connecting to kafka at %+v", cfg.Kafka.Brokers)

	return &Server{
		router: mux.NewRouter(),
		Cfg:    *cfg,
		log:    log,
	}
}

func loadVars(c *config.Config) error {
	c.Env.GetKeys()
	c.Logger.GetLoggerEnv()
	c.Server.GetHTTPSEnv()
	c.Kafka.GetKafkaEnv()
	if _, err := c.Monitoring.GetMonitoringEnv(); err != nil {
		return err
	}

	return nil
}

// Run the https server
func (s *Server) Run() error {
	s.router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)

	apiRouter := s.router.PathPrefix("/api").Subrouter()

	healthcheckSvc, _ := healthchecksvc.NewHealthCheckSvc(s.log)
	healthcheckHandler := healthcheckrest.NewHealthCheckHandlers(apiRouter, s.log, s.Cfg, healthcheckSvc, s.metricsCollector)
	healthcheckHandler.RegisterRouter()

	sqliteClient, err := sqlclient.NewSqlite()
	if err != nil {
		return err
	}
	s.log.Info("Connected to sqlite successfully")

	kafkaReqSqlRepo := kafkareqrepo.NewSqlRepo(sqliteClient, s.log)
	kafkaReqSvc := kafkareqsvc.NewKafkaRequestSvc(kafkaReqSqlRepo, s.log, s.Cfg)
	kafkReqHandler := kafkarequestres.NewKafkaRequestHandlers(apiRouter, s.log, s.Cfg, kafkaReqSvc, s.metricsCollector)
	kafkReqHandler.RegisterRouter()

	kafkaSvc := kafkasvc.NewKafkaSvc(nil, nil, nil, s.log, s.Cfg)
	kafkaHandler := kafkarest.NewKafkaHandlers(apiRouter, s.log, &s.Cfg, kafkaSvc, s.metricsCollector)
	kafkaHandler.RegisterRouter()

	runHTTP := func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Printf("Server listening on port: %s ...", s.Cfg.Server.Port)

		if err := http.ListenAndServe(fmt.Sprintf(":%s", s.Cfg.Server.Port), s.router); err != nil {
			log.Fatal("ListenAndServe error: ", err)
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go runHTTP(wg)
	wg.Wait()

	return nil
}
