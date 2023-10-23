package kafkasvc

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (svc *KafkaSvc) produceMessage(ctx context.Context, workerID int, topic string, msgs kafka.Message) error {

	if err := svc.kafkaProducer.PublishMessage(ctx, msgs); err != nil {
		return err
	}
	return nil
}

func (svc *KafkaSvc) worker(ctx context.Context, workerID int, lb chan int, topic string, msgs kafka.Message, successChan chan int, failedChan chan int) {
	for msgIdx := range lb {
		svc.log.Debugf("Message ID: %d", msgIdx)
		if err := svc.produceMessage(ctx, workerID, topic, msgs); err != nil {
			svc.log.Errorf("send message to topic: %s failed %+v", topic, err)
			failedChan <- 1
		} else {
			successChan <- 1
		}
	}
}

func (svc *KafkaSvc) SendMessage(ctx context.Context, topic string, msg string, numMessages int) (int, int) {
	numWorkers := 10

	successChan := make(chan int)
	failedChan := make(chan int)

	lb := make(chan int, numWorkers)

	msgs := kafka.Message{
		Topic: topic,
		Value: []byte(msg),
	}
	svc.log.Debugf("message: %s", msg)

	for i := 0; i < numWorkers; i++ {
		go svc.worker(ctx, i, lb, topic, msgs, successChan, failedChan)
	}

	go func() {
		for i := 0; i < numMessages; i++ {
			lb <- i
		}

		close(lb)
	}()

	successMsg := 0
	failedMsg := 0

	for i := 0; i < numMessages; i++ {
		select {
		case <-failedChan:
			failedMsg++
		case <-successChan:
			successMsg++
		}
	}
	return successMsg, failedMsg
}
