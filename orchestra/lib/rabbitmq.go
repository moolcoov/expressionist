package lib

import (
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RbMQ = &RabbitWithMutex{Mu: &sync.RWMutex{}}

	rbaddress = Getenv("RABBITMQ_ADDRESS", "localhost:5672")
	//rbchan    = Getenv("RABBITMQ_CHANNEL", "expressions")
	rbuser = Getenv("RABBITMQ_USER", "guest")
	rbpass = Getenv("RABBITMQ_PASSWORD", "guest")
)

// RabbitWithMutex Подключение и канал RabbitMQ с мьютексом для защиты от гонки
type RabbitWithMutex struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Mu      *sync.RWMutex
}

// SetupRabbit подключает RabbitMQ в первый раз
func SetupRabbit() {
	// Инициализируем
	connectionString := fmt.Sprintf("amqp://%s:%s@%s/", rbuser, rbpass, rbaddress)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Println("ERR: Failed to connect to RabbitMQ")
		log.Fatal(err.Error())
	}

	RbMQ.Conn = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Println("ERR: Failed to open a RabbitMQ channel")
		log.Fatal(err.Error())
	}

	RbMQ.Channel = ch

	fmt.Println("SETUP: RabbitMQ has been connected")
}
