package kafkamodel

type Request struct {
	Title    string `json:"title"`
	Topic    string `json:"topic,omitempty"`
	Quantity int    `json:"quantity"`
	Type     string `json:"type"` // json, text ...
	Message  string `json:"message"`
}

type ListRequestResponse struct {
	Requests []Request `json:"requests"`
}
