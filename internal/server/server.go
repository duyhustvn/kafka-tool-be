package server

import (
	"context"
	"fmt"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	healthchecksvc "kafkatool/internal/modules/healthcheck/service"
	healthcheckrest "kafkatool/internal/modules/healthcheck/transport/rest"
	kafkasvc "kafkatool/internal/modules/kafka/service"
	kafkarest "kafkatool/internal/modules/kafka/transport/res"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"

	kafkaclient "kafkatool/pkg/kafka"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type Server struct {
	router           *mux.Router
	kafkaConn        *kafka.Conn
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
	kafkaConn, err := kafkaclient.NewKafkaConnection(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Cannot connect to kafka %+v", err)
	} else {
		log.Infof("Connected to kafka")
	}

	return &Server{
		router:    mux.NewRouter(),
		Cfg:       *cfg,
		log:       log,
		kafkaConn: kafkaConn,
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
func (s *Server) Run() {
	defer s.kafkaConn.Close()

	s.router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	apiRouter := s.router.PathPrefix("/api").Subrouter()

	healthcheckSvc, _ := healthchecksvc.NewHealthCheckSvc(s.log)
	healthcheckHandler := healthcheckrest.NewHealthCheckHandlers(apiRouter, s.log, s.Cfg, healthcheckSvc, s.metricsCollector)
	healthcheckHandler.RegisterRouter()

	kafkaProducer := kafkaclient.NewProducer(s.Cfg.Kafka.Brokers, s.log)
	defer kafkaProducer.Close()

	kafkaConsumerGroup := kafkaclient.NewConsumerGroup(s.Cfg.Kafka.Brokers, s.Cfg.Kafka.GroupID, s.log)

	kafkaSvc := kafkasvc.NewKafkaSvc(s.kafkaConn, kafkaProducer, kafkaConsumerGroup, s.log, s.Cfg)
	kafkaHandler := kafkarest.NewKafkaHandlers(apiRouter, s.log, s.Cfg, kafkaSvc, s.metricsCollector)
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
}
