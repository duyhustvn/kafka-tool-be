package kafkarest

import (
	"encoding/json"
	"kafkatool/internal/common"
	kafkamodel "kafkatool/internal/modules/kafka/models"
	"net/http"
)

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
