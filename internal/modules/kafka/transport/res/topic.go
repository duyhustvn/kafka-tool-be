package kafkarest

import (
	"kafkatool/internal/common"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	"net/http"
)

func (handler *kafkaHandlers) ListTopicHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		topics, err := handler.kafkaSvc.ListTopic(ctx)
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
