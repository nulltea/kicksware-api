package events

import (
	"github.com/streadway/amqp"
)

func (b *Broker) Emmit(routingKey string, msg interface{}) error {
	data, err := b.Serializer.Encode(msg); if err != nil {
		return err
	}

	return b.Channel.Publish(
		b.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}