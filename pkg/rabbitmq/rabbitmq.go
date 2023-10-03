package rabbitmq

import (
	"MarketEye/config"
	"fmt"
	"github.com/streadway/amqp"
)

func NewRabbitMQConn(cfg *config.Config) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.Access.User,
		cfg.RabbitMQ.Access.Password,
		cfg.RabbitMQ.Access.Host,
		cfg.RabbitMQ.Access.Port,
	)
	return amqp.Dial(connAddr)
}
