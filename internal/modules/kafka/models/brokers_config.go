package kafkamodel

import (
	"errors"
	"strings"
)

type BrokersConfig struct {
	Url      string `json:"brokers_url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (bc BrokersConfig) Validator() error {
	if strings.TrimSpace(bc.Url) == "" {
		return errors.New("brokers url cannot be null")
	}
	return nil
}
