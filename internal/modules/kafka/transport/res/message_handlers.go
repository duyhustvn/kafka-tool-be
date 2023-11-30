package kafkarest

import (
	"encoding/json"
	"kafkatool/internal/common"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (handler *kafkaHandlers) CreateRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var kkrequest kafkamodel.Request
		if err := json.NewDecoder(r.Body).Decode(&kkrequest); err != nil {
			handler.log.Errorf("[CreateRequestHandler] decode request body %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		if err := kkrequest.Validator(); err != nil {
			handler.log.Errorf("[CreateRequestHandler] validation failed %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		handler.log.Debugf("kkrequest: %+v", kkrequest)

		if err := handler.kafkaSvc.CreateRequest(ctx, kkrequest); err != nil {
			handler.log.Errorf("[CreateRequestHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		common.ResponseOk(w, http.StatusOK, nil)
	}
}

func (handler *kafkaHandlers) UpdateRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestIDStr := mux.Vars(r)["request_id"]
		requestID, err := strconv.Atoi(requestIDStr)
		if err != nil {
			handler.log.Errorf("[UpdateRequestHandler] invalid request id %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		var kkrequest kafkamodel.Request
		if err := json.NewDecoder(r.Body).Decode(&kkrequest); err != nil {
			handler.log.Errorf("[UpdateRequestHandler] decode request body %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		if err := kkrequest.Validator(); err != nil {
			handler.log.Errorf("[UpdateRequestHandler] validation failed %+v", err)
			common.ResponseError(w, http.StatusBadRequest, nil, err.Error())
			return
		}

		handler.log.Debugf("kkrequest: %+v", kkrequest)

		if err := handler.kafkaSvc.UpdateRequest(ctx, requestID, kkrequest); err != nil {
			handler.log.Errorf("[UpdateRequestHandler] %+v", err)
			common.ResponseError(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		common.ResponseOk(w, http.StatusOK, nil)
	}
}
