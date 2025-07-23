package kafkamodel

type SendMessageRequestBody struct {
	Topic    string `json:"topic"`
	Message  string `json:"message"`
	Header   string `json:"header"`
	Key      string `json:"key"`
	Quantity int    `json:"quantity"` // number of message to send
}

type SendMessageResponse struct {
	TotalMessage int `json:"totalMessage"`
	Success      int `json:"success"`
	Failed       int `json:"failed"`
}
