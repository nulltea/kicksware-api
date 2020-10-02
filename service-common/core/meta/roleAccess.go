package meta

import "go.kicksware.com/api/user-service/core/model"

var (
	ZeroAccess = []model.UserRole{}
	GuestAccess = []model.UserRole{ model.Guest }
	RegularAccess = []model.UserRole{ model.Guest, model.Regular, model.Admin }
	UserAccess = []model.UserRole{ model.Regular, model.Admin }
	AdminAccess = []model.UserRole{ model.Admin }
)