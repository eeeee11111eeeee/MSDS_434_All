package main

import (
        "fmt"
        "github.com/streadway/amqp"
)


func main() {
        fmt.Println("Go RabbitMQ Tutorial")

        conn, err := amqp.Dial("amqp://guest:guest@ec2-35-166-104-128.us-west-2.compute.amazonaws.com:5672/")
        if err != nil {
                fmt.Println(err)
                panic(err)
        }
        defer conn.Close()
        fmt.Println("Successfully Connected to ou RabbitMQ instance")

        ch,err := conn.Channel()
        if err != nil {
                fmt.Println(err)
                panic(err)
        }

	// test queue declaration
	// name
	// durable
	// autoDelete
	// exclusive
	// noWait
	// args.
        defer ch.Close()
        q, err := ch.QueueDeclare(
                "TestQueue",
                false,
                false,
                false,
                false,
                nil,
        )

        if err != nil {
                fmt.Println(err)
                panic(err)
        }
        fmt.Println(q)


       err= ch.Publish(
                "",
                "TestQueue",
                false,
                false,
                amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte("New Hello World"),
                },
        )

        if err != nil {
                fmt.Println(err)
                panic(err)
        }

        fmt.Println("Successfully published message to Queue")




}
