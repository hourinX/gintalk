package systems

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func InitRabbitMQ() error {
	conf := GetConfig().RabbitMQ
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.VHost,
	)

	conn, err := amqp.Dial(uri)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	RabbitMQConn = conn
	RabbitMQChannel = ch
	fmt.Println("✅ RabbitMQ链接成功")
	return nil
}

func CloseRabbitMQ() {
	if RabbitMQChannel != nil {
		_ = RabbitMQChannel.Close()
	}
	if RabbitMQConn != nil {
		_ = RabbitMQConn.Close()
	}
}
