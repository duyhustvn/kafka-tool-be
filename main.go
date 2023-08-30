package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(ctx context.Context, brokerUrl []string, topic string, data string, idx int, successChan chan int, failChan chan int) {
	log.Printf("Start sending message to topic: %s", topic)
	w := &kafka.Writer{
		Addr:  kafka.TCP(brokerUrl...),
		Topic: topic,
	}

	if err := w.WriteMessages(ctx, kafka.Message{
		Value: []byte(data),
	}); err != nil {
		log.Printf("There is error produce message %+v", err)
		failChan <- idx
	}

	successChan <- idx
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

		// Break until \n
		break
	}
	lines = strings.TrimSpace(lines)
	return lines
}

func worker(ctx context.Context, workers chan int, brokers []string, topic string, data string, successChan chan int, failChan chan int) {
	for idx := range workers {
		ProduceMessage(ctx, brokers, topic, data, idx, successChan, failChan)
	}
}

// produce n messages
func ProduceMessages(ctx context.Context, brokers []string, topic string, message string, n int) Response {
	workers := make(chan int, 1000)

	successChan := make(chan int)
	failChan := make(chan int)

	for i := 0; i < cap(workers); i++ {
		go worker(ctx, workers, brokers, topic, message, successChan, failChan)
	}

	go func() {
		for i := 0; i <= n; i++ {
			workers <- i
		}
	}()

	successCounter := 0
	failedCounter := 0
	for i := 0; i < n; i++ {
		select {
		case <-successChan:
			successCounter++
		case <-failChan:
			failedCounter++
		}
	}

	response := Response{
		TotalMessage: n,
		Success:      successCounter,
		Failed:       failedCounter,
	}

	return response
}

func cmd() {
	brokerUrl := StringPrompt("Enter broker url (Left empty for default value: 172.17.0.1:9092) >>")
	if brokerUrl == "" {
		brokerUrl = "172.17.0.1:9092"
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
			invalidNumber := true
			for invalidNumber {
				message_nums_str := StringPrompt("Enter the number of messages >>")
				message_num, err := strconv.Atoi(message_nums_str)
				if err != nil {
					fmt.Println("Invalid number")
				}
				ProduceMessages(context.Background(), []string{brokerUrl}, topic, produceMessage, message_num)
				invalidNumber = false
			}
		}
		produceMessage = ""
	}
}

type RequestBody struct {
	Topic    string `json:"topic"`
	Message  string `json:"message"`
	Quantity int    `json:"quantity"`
}

type Response struct {
	TotalMessage int `json:"totalMessage"`
	Success      int `json:"success"`
	Failed       int `json:"failed"`
}

func PublishMessageHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read body: %s\n", err)
		return
	}

	var requestBody RequestBody
	if err := json.Unmarshal(body, &requestBody); err != nil {
		log.Printf("Unmarshal body failed %+v", err)
		return
	}

	res := ProduceMessages(context.Background(), []string{"172.17.0.1:9092"}, requestBody.Topic, requestBody.Message, requestBody.Quantity)

	b, err := json.Marshal(res)
	if err != nil {
		log.Printf("Marshal body failed %+v", err)
		return
	}

	io.WriteString(w, string(b))
}

func ListTopic() []string {
	topics := []string{}
	conn, err := kafka.Dial("tcp", "172.17.0.1:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	for k := range m {
		topics = append(topics, k)
	}
	return topics
}

func main() {
	log.Println("Start program")
	go func() {
		topics := ListTopic()
		log.Println("List topic ", topics)
	}()
	http.HandleFunc("/", PublishMessageHandler)

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalln("Failed to listen on port 9000 ", err)
	}
}
