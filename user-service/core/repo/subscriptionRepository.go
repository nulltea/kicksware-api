package repo

import "go.kicksware.com/api/user-service/core/model"

type SubscriptionRepository interface {
	Add(record model.MailSubscription) error
	Delete(email string) error
}
