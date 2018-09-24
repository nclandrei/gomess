package main

import (
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
	"sync"
	"time"
)

const (
	topic = "SVERIGE"
)

func main() {
	var wg sync.WaitGroup
	var inc int
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			producer.Publish(topic, []byte(fmt.Sprintf("message #%d", inc)))
			inc++
			fmt.Println("message sent...")
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
