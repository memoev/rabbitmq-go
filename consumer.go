package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer connected!")

	conn, err := amqp.Dial("amqps://kwnwkymm:aCdrri4zkPRxoWsoKxPd62VndTEifNea@gull.rmq.cloudamqp.com/kwnwkymm")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"Test Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Connected to RabbitMq instance")
	fmt.Println(" [*] - waiting for msgs")
	<-forever
}
