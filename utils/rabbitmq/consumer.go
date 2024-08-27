package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-libr-category/utils/constant"
	"go-libr-category/utils/email"
	"go-libr-category/utils/logger"
)

func (r *RabbitMQ) Consume() (err error) {
	//forever := make(chan bool)
	r.ConsumeEmailQueue()

	//<-forever

	return
}

func (r *RabbitMQ) ConsumeEmailQueue() {
	var (
		queue        = constant.EmailQueue
		exchangeName = viper.GetString("name")
		routingKey   = viper.GetString("app.mode")
	)

	emailQueue := MqConfig{
		QueueName:    queue,
		RoutingKey:   routingKey,
		ExchangeName: exchangeName,
	}

	_ = r.DeclareExchange(emailQueue)
	_ = r.DeclareQueue(emailQueue)
	_ = r.Bind(emailQueue)

	// declaring consumer with its properties over channel opened
	msgEmailQueues, err := r.Channel.Consume(
		emailQueue.QueueName.String(), // queue name
		"",                            // consumer
		true,                          // auto ack
		false,                         // exclusive
		false,                         // no local
		false,                         // no wait
		nil,                           //args
	)
	if err != nil {
		panic(err)
	}

	// do your logic here
	go func() {
		for msg := range msgEmailQueues {
			ctx := &gin.Context{}

			var emailNotif email.EmailNotif
			err = json.Unmarshal(msg.Body, &emailNotif)
			if err != nil {
				logger.ErrorWithCtx(ctx, nil, err)
			}

			err = emailNotif.SendEmail()
			if err != nil {
				logger.ErrorWithCtx(ctx, nil, err)
				fmt.Printf("Received Error Message: From ConsumeEmailQueue \n")
				return
			}

			fmt.Printf("Received Success Message: From ConsumeEmailQueue \n")
		}
	}()
}
