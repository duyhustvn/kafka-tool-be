package kafkamodel

import (
	"errors"
	"strings"
)

type BrokersConfig struct {
	Url string `json:"brokers_url"`
}

func (bc BrokersConfig) Validator() error {
	if strings.TrimSpace(bc.Url) == "" {
		return errors.New("brokers url cannot be null")
	}
	return nil
}
