package factory

import (
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/rating/api/events"
	"go.kicksware.com/api/services/rating/core/service"
	"go.kicksware.com/api/services/rating/env"
)

func ProvideEventBus(service service.RatingService, auth core.AuthService, config env.ServiceConfig) *events.Handler {
	return events.NewHandler(service, config.EventBus)
}
