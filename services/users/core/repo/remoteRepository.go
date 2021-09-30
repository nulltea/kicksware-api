package repo

import "go.kicksware.com/api/services/users/core/model"

type RemoteRepository interface {
	Connect(userID string, remoteID string, provider model.UserProvider) error
	Sync(userID string, remotes map[model.UserProvider]string) error
	Track(remoteID string, provider model.UserProvider) (string, error)
}
