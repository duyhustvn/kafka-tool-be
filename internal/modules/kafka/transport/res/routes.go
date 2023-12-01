package kafkarest

import "net/http"

func (handler *kafkaHandlers) RegisterRouter() {
	s := handler.router.PathPrefix("/kafka").Subrouter()
	// connect kafka brokers
	s.HandleFunc("/brokers/connect", handler.ConnectKafkaBrokerHandler()).Methods(http.MethodPost)
	// Topic
	s.HandleFunc("/topics", handler.ListTopicHandler()).Methods(http.MethodGet)
	// Publish
	s.HandleFunc("/publish", handler.SendMessageHandler()).Methods(http.MethodPost)
	// Subscribe
	s.HandleFunc("/subscribe/topic/{topic_name}", handler.SubscriberHandler()).Methods(http.MethodGet)
}
