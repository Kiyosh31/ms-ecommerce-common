package messaging

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// func ConnectAmqp(user, pass, host, port string) (*amqp.Channel, func() error) {
// address := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
func ConnectAmqp(host string) (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s", host)

	conn, err := amqp.Dial(address)
	if err != nil {
		log.Fatalf(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = channel.ExchangeDeclare("exchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return channel, conn.Close
}
