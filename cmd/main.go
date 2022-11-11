package main

import (
	"RestApiLibrary/internal/config"
	server "RestApiLibrary/internal/server"
	"RestApiLibrary/pkg/database"
	"log"
)

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name Sagindykov Arsan
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	c := new(config.Config)
	c, err := config.LoadConfig()
	db, err := database.ConnectToPostgresDB(c)
	if err != nil {
		log.Fatalln(err)
	}

	server := server.NewServer(db)

	server.ServerRun(c)
}
