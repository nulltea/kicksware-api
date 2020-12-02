package events

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func (bus *EventBus) Publish(queue, routingKey string, msg interface{}) error {
	data, err := json.Marshal(msg); if err != nil {
		return err
	}

	return bus.Channel.Publish(
		bus.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}