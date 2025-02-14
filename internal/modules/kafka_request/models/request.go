package kafkareqmodel

type Request struct {
	ID       string `json:"id,omitempty"`
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

type CreateNewRequestResponse struct {
	NewRequestId int64 `json:"new_request_id"`
}
