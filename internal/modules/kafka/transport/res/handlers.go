package kafkarest

import (
	"encoding/json"
	"kafkatool/internal/common"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"kafkatool/internal/metrics"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	kafkasvc "kafkatool/internal/modules/kafka/service"
	"net/http"

	"github.com/gorilla/mux"
)

type kafkaHandlers struct {
	router           *mux.Router
	kafkaSvc         kafkasvc.KafkaSvc
	log              logger.Logger
	cfg              config.Config
	metricsCollector metrics.IMetricCollector
}

func NewKafkaHandlers(router *mux.Router, log logger.Logger, cfg config.Config, kafkaSvc *kafkasvc.KafkaSvc, metricCollector metrics.IMetricCollector) *kafkaHandlers {
	return &kafkaHandlers{router: router, log: log, cfg: cfg, kafkaSvc: *kafkaSvc, metricsCollector: metricCollector}
}

func (handler *kafkaHandlers) ListTopicHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topics, err := handler.kafkaSvc.ListTopic()
		if err != nil {
			handler.log.Errorf("[ListTopicHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}
		res := kafkamodel.ListTopicResponse{
			Topics: topics,
		}

		common.ResponseOk(w, http.StatusOK, res)
	}
}

func (handler *kafkaHandlers) SendMessageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body kafkamodel.SendMessageRequestBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			handler.log.Errorf("[SendMessageHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		handler.log.Debugf("[SendMessageHandler] Start sending message")
		successMsg, failedMsg := handler.kafkaSvc.SendMessage(ctx, body.Topic, body.Message, body.Quantity)

		res := kafkamodel.SendMessageResponse{
			TotalMessage: body.Quantity,
			Success:      successMsg,
			Failed:       failedMsg,
		}

		common.ResponseOk(w, http.StatusOK, res)
	}
}
