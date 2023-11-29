package kafkarest

import (
	"encoding/json"
	"kafkatool/internal/common"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	"net/http"
)

func (handler *kafkaHandlers) ListRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requests, err := handler.kafkaSvc.ListRequest(ctx)
		if err != nil {
			handler.log.Errorf("[ListTopicHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}
		res := kafkamodel.ListRequestResponse{
			Requests: requests,
		}

		common.ResponseOk(w, http.StatusOK, res)
	}
}

func (handler *kafkaHandlers) SaveRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var kkrequest kafkamodel.Request
		if err := json.NewDecoder(r.Body).Decode(&kkrequest); err != nil {
			handler.log.Errorf("[SaveRequestHandler] %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		if err := kkrequest.Validator(); err != nil {
			handler.log.Errorf("[SaveRequestHandler] validation failed %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		handler.log.Debugf("kkrequest: %+v", kkrequest)

		if err := handler.kafkaSvc.SaveRequest(ctx, kkrequest); err != nil {
			handler.log.Errorf("[SaveRequestHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		common.ResponseOk(w, http.StatusOK, nil)
	}
}
