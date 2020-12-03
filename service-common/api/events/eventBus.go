package events

import (
	"github.com/golang/glog"
	"github.com/streadway/amqp"

	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/util"
)

type EventBus struct {
	Channel *amqp.Channel
	Exchange string
}

func NewEventBus(config config.ConnectionConfig, exchange string) *EventBus {
	conn, err := amqp.DialTLS(config.Endpoint, util.NewTLSConfig(config.TLS)); if err != nil {
		glog.Fatal(err)
	}
	ch, err := conn.Channel(); if err != nil {
		glog.Fatal(err)
	}
	return &EventBus{
		Channel: ch,
		Exchange: exchange,
	}
}