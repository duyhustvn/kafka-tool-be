package kafkarest

import (
	"context"
	"fmt"
	"kafkatool/internal/common"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

func kafkaMessageProcessor(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	fmt.Println("kafkaMessageProcessor")
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("kafkaMessageProcessor Done")
			return
		default:
		}
		fmt.Println("fetchMessage")

		msg, err := r.FetchMessage(ctx)
		if err != nil {
			// s.log.Errorf("[ProcessMessage] FetchMessage workerID: %d, err: %+v", workerID, err)
			fmt.Printf("[ProcessMessage] FetchMessage workerID: %d, err: %+v", workerID, err)
			continue
		}

		fmt.Println(string(msg.Value))
	}
}

func (handler *kafkaHandlers) SubscriberHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		topic := mux.Vars(r)["topic_name"]
		if strings.TrimSpace(topic) == "" {
			common.ResponseError(w, http.StatusBadRequest, nil, "topic cannot be nil")
		}
		handler.log.Infof("Subscribe to topic: %+v", topic)
		handler.kafkaSvc.ConsumeMessage(ctx, []string{topic}, 1, kafkaMessageProcessor)

		common.ResponseOk(w, http.StatusOK, nil)
	}
}
