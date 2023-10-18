package kafkarest

import "net/http"

func (handler *kafkaHandlers) RegisterRouter() {
	s := handler.router.PathPrefix("/kafka").Subrouter()

	s.HandleFunc("/topics", handler.ListTopicHandler()).Methods(http.MethodGet)
	s.HandleFunc("/messages", handler.SendMessageHandler()).Methods(http.MethodPost)
}
