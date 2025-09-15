package kafkaclient

import (
	"context"
	"kafkatool/internal/config"
	"kafkatool/internal/logger"
	"sync"

	"github.com/segmentio/kafka-go"
)

// MessageProcessor processor methods must implement kafka.Worker func method interface
type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

// Worker kafka consumer worker fetch and process messages from reader
type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type ConsumerGroup interface {
	// ConsumeTopic(ctx context.Context, cancel context.CancelFunc, groupID, topic string, poolSize int, worker Worker)
	// GetNewKafkaReader(kafkaURL []string, topic, groupID string) *kafka.Reader
	// GetNewKafkaWriter(topic string) *kafka.Writer

	ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker)
}

type consumerGroup struct {
	kafkaCfg config.Kafka
	GroupID  string
	log      logger.Logger
}

// NewConsumerGroup kafka consumer group constructor
func NewConsumerGroup(
	kafkaCfg config.Kafka,
	groupID string,
	log logger.Logger,
) *consumerGroup {
	return &consumerGroup{
		kafkaCfg: kafkaCfg,
		GroupID:  groupID,
		log:      log,
	}
}

// ConsumeTopic start consumer group with given worker and pool size
func (c *consumerGroup) ConsumeTopic(
	ctx context.Context,
	groupTopics []string,
	poolSize int,
	worker Worker,
) {
	r := GetNewKafkaReader(c.kafkaCfg, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Warnf("consumerGroup.r.Close: %v", err)
		}
	}()

	c.log.Infof("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
