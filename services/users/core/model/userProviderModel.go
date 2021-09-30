package model

type UserProvider string

var (
	Internal UserProvider = ""
	Facebook UserProvider = "facebook"
	Google   UserProvider = "google"
	GitHub   UserProvider = "github"
	Apple    UserProvider = "apple"
)
