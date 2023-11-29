package kafkarest

import (
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
