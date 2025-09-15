package kafkarest

import (
	"encoding/json"
	"kafkatool/internal/common"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	"net/http"
)

func (handler *kafkaHandlers) ConnectKafkaBrokerHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var body kafkamodel.BrokersConfig
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			handler.log.Errorf("[UpdateBrokerHandler] %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		if err := body.Validator(); err != nil {
			handler.log.Errorf("[UpdateBrokerHandler] data validator failed %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		if err := handler.kafkaSvc.ConnectKafkaBrokers(ctx, body); err != nil {
			handler.log.Errorf("[UpdateBrokerHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		common.ResponseOk(w, http.StatusOK, nil)
	}
}
