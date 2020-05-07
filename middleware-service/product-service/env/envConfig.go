package env

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
)

var (
	ProjectDirectory, _ = os.Getwd()
	Environment = os.Getenv("ENV")
	Host = os.Getenv("HOST")
	HostName = os.Getenv("HOSTNAME")
	ServiceConfigPath = ProjectDirectory + os.Getenv("CONFIG_PATH")
)

func InitEnvironment() {
	if os.Getenv("DEBUG") == "True" {
		err := env.Load(ProjectDirectory + "/env/.env"); if err != nil {
			log.Fatal(err)
		}
		reassignVariables()
	}
}

func reassignVariables() {
	Environment = os.Getenv("ENV")
	Host = os.Getenv("HOST")
	HostName = os.Getenv("HOSTNAME")
	ServiceConfigPath = ProjectDirectory + os.Getenv("CONFIG_PATH")
}