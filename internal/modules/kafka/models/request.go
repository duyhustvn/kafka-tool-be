package kafkamodel

type Request struct {
	Title    string `json:"title"`
	Topic    string `json:"topic,omitempty"`
	Quantity uint   `json:"quantity"`
	Type     string `json:"type"` // json, text ...
	Message  string `json:"message"`
}

func (r Request) Validator() error {
	return nil
}

type ListRequestResponse struct {
	Requests []Request `json:"requests"`
}
