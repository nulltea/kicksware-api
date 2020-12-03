package factory

import (
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/rating-service/api/events"
	"go.kicksware.com/api/rating-service/core/service"
	"go.kicksware.com/api/rating-service/env"
)

func ProvideEventBus(service service.RatingService, auth core.AuthService, config env.ServiceConfig) *events.Handler {
	return events.NewHandler(service, config.EventBus)
}