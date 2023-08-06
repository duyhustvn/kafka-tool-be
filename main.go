package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(ctx context.Context, brokerUrl []string, topic string, data string) {
	w := &kafka.Writer{
		Addr:  kafka.TCP(brokerUrl...),
		Topic: topic,
	}

	if err := w.WriteMessages(ctx, kafka.Message{
		Value: []byte(data),
	}); err != nil {
		log.Printf("There is error produce message %+v", err)
		return
	}

	log.Println("Send message successfully")
}

func StringPrompt(label string) string {
	fmt.Println(label)
	reader := bufio.NewReader(os.Stdin)

	lines := ""
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		//append the line to a slice
		lines += line
	}
	lines = strings.TrimSpace(lines)
	return lines
}

func main() {
	brokerUrl := StringPrompt("Enter broker url (Left empty for default value: localhost:9092) >>")
	if brokerUrl == "" {
		brokerUrl = "localhost:9092"
	}
	fmt.Println("brokerUrl: ", brokerUrl)

	topic := StringPrompt("Enter the topic name >>")
	for topic == "" {
		topic = StringPrompt("Enter the topic name >>")
	}
	fmt.Println("topic: ", topic)

	produceMessage := ""
	for produceMessage == "" {
		produceMessage = StringPrompt("Enter produce Message >>")
		if produceMessage != "" {
			ProduceMessage(context.Background(), []string{brokerUrl}, topic, produceMessage)
		}
		produceMessage = ""
	}
}
