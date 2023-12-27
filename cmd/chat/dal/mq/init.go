package mq

import (
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

var (
	Rmq *RabbitMQ
	Mu  sync.Mutex
)

func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		mqurl: utils.GetMQUrl(),
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	if err != nil {
		klog.Error(err)
		return
	}
	Rmq.conn = dial
}

func (r *RabbitMQ) destroy() {
	r.conn.Close()
}
