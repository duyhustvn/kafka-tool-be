package kafkarequestres

import "net/http"

func (handler *kafkaHandlers) RegisterRouter() {
	s := handler.router.PathPrefix("/kafka").Subrouter()
	// Request
	s.HandleFunc("/requests", handler.ListRequestHandler()).Methods(http.MethodGet)
	s.HandleFunc("/requests", handler.CreateRequestHandler()).Methods(http.MethodPost)
	s.HandleFunc("/requests/{request_id}", handler.UpdateRequestHandler()).Methods(http.MethodPut)
}
