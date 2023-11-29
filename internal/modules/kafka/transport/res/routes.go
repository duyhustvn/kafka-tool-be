package kafkarest

import "net/http"

func (handler *kafkaHandlers) RegisterRouter() {
	s := handler.router.PathPrefix("/kafka").Subrouter()

	s.HandleFunc("/topics", handler.ListTopicHandler()).Methods(http.MethodGet)
	s.HandleFunc("/publish", handler.SendMessageHandler()).Methods(http.MethodPost)
	s.HandleFunc("/subscribe/topic/{topicName}", handler.SubscriberHandler()).Methods(http.MethodGet)
	s.HandleFunc("/requests", handler.ListRequestHandler()).Methods(http.MethodGet)
}
