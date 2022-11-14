package main

import (
	_ "RestApiLibrary/docs"
	server "RestApiLibrary/internal/server"
	"RestApiLibrary/pkg/config"
	"RestApiLibrary/pkg/database"
	"log"
)

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  Sagindykov Arsan
// @contact.email ArsanSa@halykbank.kz

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host rest-api-library
func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := database.ConnectToPostgresDB(c)
	if err != nil {
		log.Println(err.Error())
		return
	}

	server := server.NewServer(db)

	server.ServerRun(c)
}
