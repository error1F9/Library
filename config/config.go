package config

import (
	"os"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
type Server struct {
	Port string
}

type Token struct {
	Key string
}

type AppConfig struct {
	DB     DB
	Server Server
	Token  Token
}

func NewConfig() AppConfig {
	return AppConfig{
		DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Server{
			Port: os.Getenv("SERVER_PORT"),
		},
		Token{
			Key: os.Getenv("TOKEN_KEY"),
		},
	}
}
