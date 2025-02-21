package main

import (
	"Library/config"
	_ "Library/internal/pkg/docs"
	"Library/run"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
)

//	@title			Library API
//	@version		1.0
//	@description	This is a sample Library server

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the JWT token.

//	@contact.name	Kirill Efremenko
//	@contact.email	kirikozavrr@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/library

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg := config.NewConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logger.Sync()

	app := run.NewApp(cfg, logger)
	app.Bootstrap().Run()
	os.Exit(0)
}
