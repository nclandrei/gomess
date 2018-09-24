package main

import (
	"fmt"
	"github.com/bitly/go-nsq"
	"log"
	"sync"
)

func main() {
	consumer, err := nsq.NewConsumer("SVERIGE", "ch", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body))
		return nil
	}))
	consumer.ConnectToNSQD(":4150")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
		}
	}()
	wg.Wait()
}
