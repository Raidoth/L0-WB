package config

import "log"

const (
	ViewPathMainPage = "service/view/index.html"
	ServerPort       = ":8080"
	HostPg           = "localhost"
	PortPg           = 5432
	UserPg           = "postgres"
	PasswordPg       = "qwerty"
	NameDbPg         = "WB"
)

func InfoConfigParams() {
	log.Printf("[CONFIG PARAMS]\n\tViewPathMainPage: %s\n\tServerPort: %s\n\tHostPostgres: %s\n\tPortPostgres: %d\n\tUserPostgres: %s\n\tPasswordPostgres: %s\n\tNameDatabasePostgres: %s\n", ViewPathMainPage, ServerPort[1:], HostPg, PortPg, UserPg, PasswordPg, NameDbPg)
}
