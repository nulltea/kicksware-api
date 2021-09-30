package events

import (
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/streadway/amqp"
	"go.kicksware.com/api/shared/api/events"
	"go.kicksware.com/api/shared/config"

	"go.kicksware.com/api/services/rating/core/service"
	ratingJson "go.kicksware.com/api/services/rating/usecase/serializer/json"
)

type Handler struct {
	EventBus *events.Broker
	Service  service.RatingService
}

func NewHandler(service service.RatingService, config config.ConnectionConfig) *Handler {
	return &Handler{
		EventBus: events.NewEventsBroker(config, "amq.topic", ratingJson.NewSerializer()),
		Service: service,
	}
}

func (h *Handler) SubscribeHandlers() {
	forever := make(chan bool)
	if err := h.EventBus.Consume("rating.view", "rating.view", h.viewsHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.EventBus.Consume("rating.search", "rating.search", h.searchesHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	if err := h.EventBus.Consume("rating.order", "rating.order", h.ordersHandler, 1); err != nil {
		glog.Fatalln(err)
	}
	fmt.Println("Event listeners active...")
	<- forever
}

func (h *Handler) viewsHandler(msg amqp.Delivery) bool {
	entity, ok := getEntity(msg.Body); if !ok {
		return false
	}
	if _, err := h.Service.IncrementViews(entity); err != nil {
		return false
	}
	fmt.Printf("view event handled for: %q\n", entity)
	return true
}

func (h *Handler) ordersHandler(msg amqp.Delivery) bool {
	entity, ok := getEntity(msg.Body); if !ok {
		return false
	}
	if _, err := h.Service.IncrementOrders(entity); err != nil {
		return false
	}
	fmt.Printf("order event handled for: %q\n", entity)
	return true
}

func (h *Handler) searchesHandler(msg amqp.Delivery) bool {
	entity, ok := getEntity(msg.Body); if !ok {
		return false
	}
	if _, err := h.Service.IncrementSearches(entity); err != nil {
		return false
	}
	fmt.Printf("search event handled for: %q\n", entity)
	return true
}

func getEntity(body []byte) (string, bool) {
	var entity string
	if err := json.Unmarshal(body, &entity); err != nil {
		glog.Errorln(err)
		return "", false
	}
	return entity, true
}
