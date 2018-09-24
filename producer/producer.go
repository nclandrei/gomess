package main

import (
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
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(":20497", config)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			producer.Publish(topic, []byte(`this is a test`))
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
