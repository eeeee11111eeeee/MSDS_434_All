package main


import (
        "fmt"
        "github.com/streadway/amqp"
)


func main(){
        fmt.Println("Consumer application")
        conn, err := amqp.Dial("amqp://guest:guest@ec2-35-166-104-128.us-west-2.compute.amazonaws.com:5672")
        if err != nil {
                fmt.Println(err)
                panic(err)
        }
        defer conn.Close()

        ch,err := conn.Channel()
        if err != nil {
                fmt.Println(err)
                panic(err)
        }

        defer ch.Close()

        msgs,err := ch.Consume(
                "TestQueue",
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
                        fmt.Printf("Received Messages: %s\n", d.Body)
                }
        }()

        fmt.Println("Successfully connected to our RabbitMQ instance")
        fmt.Println(" [*] - waiting for message")
        <-forever

}
